# Go for-range 循环的快照机制

调研日期：2026-09-13

## 结论

`for i, opt := range x` 中的 range 表达式 `x` 只求值一次，快照的是**进入循环那一刻 `x` 的值**（对 slice 而言是 header：指向底层数组的指针 + len + cap），它保存在编译器生成的隐藏临时量里。循环变量 `i`、`opt` 没有快照：每轮迭代重新获得值，`opt` 是从快照 header 指向的底层数组第 i 位**现读**出来的。

由此产生两条看似矛盾、实则各管一摊的规则：

- 循环体内给被遍历变量重新赋值（换 header）：循环**看不见**，迭代计划不变；
- 循环体内原地改底层数组元素：循环**看得见**，后续轮次读到的 opt 是新值。

本仓库 `backtrack/permute_test.go` 中 permute2 的 panic（`slice bounds out of range [3:2]`）正是「迭代计划由快照驱动、循环体语句读变量当前值」两者脱节的实例。

## 快照的是什么

Go 规范原文：**The range expression x is evaluated once before beginning the loop.**（range 表达式在循环开始前只求值一次。）

- 快照对象是 **range 表达式的值**。对 slice 而言就是 header（ptr + len + cap）——注意快照的**不是元素的一次性拷贝**；
- 迭代次数由快照的 len 定死：len=3 就跑 i=0,1,2 三轮；
- 循环变量每轮重新获得值：`opt := _snapshot[i]`，元素是**每轮现读**的，不是提前拷好的；
- 例外：当只有一个循环变量且 `len(x)` 是常量（如数组字面量 `for i := range [3]int{...}`）时，表达式不求值（规范原文，实践中影响甚微）；
- 附带：自 Go 1.22 起，循环变量每轮迭代是全新变量（此前是同一变量反复赋值），对本主题的结论没有影响。

## 快照存在哪里

编译器生成的**隐藏临时量**里——源码里没有它的名字，开发者无法访问，类似按值传参的形参。编译器对

```go
for i, opt := range options {
    // 循环体
}
```

的实际生成大致是：

```go
_snapshot := options              // 隐藏：进入循环时把 header 拷走
for _i := 0; _i < len(_snapshot); _i++ {
    i := _i
    opt := _snapshot[_i]          // 从快照现读元素
    // 循环体原样展开，其中的 options 仍是源码里的那个变量
}
```

所以运行时存在「两套 options」：迭代机制只认 `_snapshot`，循环体只认变量 `options`，互不干扰——这就是「只有一个源码变量，却有两个存储位置」的由来。

## 影响：可见性一览

| 循环体内做的事 | 循环看得见吗 | 原因 |
|---|---|---|
| 给被遍历变量重新赋值（换 header） | 看不见 | 快照 header 是独立副本 |
| 原地改元素（如 `options[0] = 9`） | 看得见 | 共享底层数组，元素每轮现读 |
| 修改循环变量 `i`、`opt` | 看不见 | 下一轮开始会被重新赋值覆盖 |

## 代码示例

以下示例均可在本地 `go run` 验证，输出为实际运行结果。

### 示例 1：换 header，循环看不见

```go
x := []int{1, 2, 3}
for i, v := range x {
    if i == 1 {
        x = []int{9, 8, 7} // 给变量换 header：迭代计划不受影响
    }
    fmt.Println(i, v)
}
// 输出：
// 0 1
// 1 2
// 2 3        ← 仍按进入时的 [1,2,3] 走完三轮
```

### 示例 2：原地改元素，循环看得见

```go
x := []int{1, 2, 3}
for i, v := range x {
    if i == 0 {
        x[2] = 99 // 原地改：共享底层数组
    }
    fmt.Println(i, v)
}
// 输出：
// 0 1
// 1 2
// 2 99       ← 第 3 轮的 v 是现读出来的，看到了修改
```

### 示例 3：append 到被遍历的变量，新元素不会被迭代

```go
x := []int{1, 2, 3}
for _, v := range x {
    x = append(x, v*10) // 迭代长度进入循环时就定死
}
fmt.Println(x)
// 输出：[1 2 3 10 20 30]
// 只跑了 3 轮——新 append 的元素不参与本轮迭代，否则会无限循环
```

### 示例 4：「快照 vs 变量」脱节导致越界（permute2 的最小复现）

```go
x := []int{1, 2, 3}
for i, v := range x {
    if i+1 == len(x) { // len(x) 读的是变量当前值
        x = nil
    } else {
        x = x[i+1:]
    }
    x = append([]int{v}, x...) // 恢复：i>0 时丢掉前缀
}
// panic: runtime error: slice bounds out of range [3:2]
// 第 3 轮：i=2 由快照驱动照常到来，但变量 x 已被上一轮恢复成 [2,3]，
// x[3:] 起点 3 超长 → 越界。
```

## 对开发者的影响

1. **不要在 range 体内给被遍历的 slice 变量重新赋值**。否则 `i`（由快照驱动）与 `len(x)`（读变量当前值）会脱节，轻则逻辑错误，重则越界 panic——permute2 的 `[3:2]` 就是这么来的（最小复现见示例 4）。
2. **不要在 range 体内对同一个 slice 变量 append 并指望新元素被迭代**。迭代长度在进入循环时就定死：未扩容时新元素不参与本轮迭代；扩容重分配后循环仍迭代旧数组。需要动态增删时，改用索引循环 `for i := 0; i < len(x); i++` 或先拷贝一份（见示例 3）。
3. **原地改元素是安全的**，且对后续轮次可见。
4. 其他类型的求值语义：数组与 slice 相同；字符串先解析一次再迭代（双变量形式按 UTF-8 解码为 rune）；map 对 map 值求值一次，迭代期间增删元素对后续迭代的影响没有保证，不要依赖；channel 的接收绑定一次，迭代到 channel 关闭为止。

## 参考

- [Go 规范：For statements with range clause](https://go.dev/ref/spec#For_statements)
- 本仓库案例：`backtrack/permute_test.go` 中 permute2 的调用树注释（错点1～4 与越界推演）
