package backtrack

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/VvJkup/description/
// 对比 backtrack/subsets_test.go
// 更推荐这个版本
func permute(nums []int) [][]int {
	result := make([][]int, 0)

	// 决策阶段：集合元素中的每个元素对应一个阶段
	// 选择列表：因为是排列，所以处理 nums[k] 只有一种选择：添加进 path。
	//         且需要更改下一阶段的选择列表，即将已处理的元素从选择列表中移除。
	//         这里选择不用单独的对象表示选择列表，从 nums 中动态过滤出选择列表
	// 路径：保存已添加到当前排列中的元素
	var backtrack func(nums []int, k int, path []int)
	backtrack = func(nums []int, k int, path []int) {
		// 结束条件
		if k == len(nums) {
			result = append(result, append([]int(nil), path...))
			return
		}

		// 对于一个具体排列，假设已选出 k 个元素，
		// 那么对于第 k+1 个就有 n-k 个选择，每一种选择都需要遍历
		// 由于这个实现选择不用单独的对象表示选择列表，因此需要遍历
		// 整个 nums，然后动态判断是否是可用选择
		for _, item := range nums {
			if slices.Contains(path, item) { // 是否是不可用选择
				continue
			}

			path = append(path, item)
			backtrack(nums, k+1, path)
			// 恢复
			path = path[:len(path)-1]
		}
	}

	backtrack(nums, 0, []int{})
	return result
}

func permute2(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func(int, []int, []int)
	// 不能像下面这样声明与定义一行语句搞定，因为函数内部会递归调用其自身，
	// 导致报告：undefined: backtrack。必须先声明，再定义。
	// backtrack := func(k int, options []int, path []int) {
	backtrack = func(k int, options []int, path []int) {
		if k == len(nums) {
			// path 虽然是值传递（slice header 按值复制），但底层数组是共享的：
			// 初始 cap=3，整棵递归树上所有分支的 append 都复用同一块数组，
			// 后续分支会继续改写它。所以这里必须拷贝一份，再存入 result。
			result = append(result, append([]int(nil), path...))
			// 拷贝后，DST 为空列表，不是期望结果
			// dst := make([]int, 0, len(path))
			// dst := make([]int, len(path)) // 这样才 OK
			// _ = copy(dst, path)
			// result = append(result, dst)
			return
		}

		for _, opt := range options {
			// 这里选择不更改选择列表，在每次循环开始时，动态过滤出当前阶段对应的选择
			if slices.Contains(path, opt) { // Go 1.21 之前标准库并不支持
				continue
			}
			path = append(path, opt)
			backtrack(k+1, options, path)
			// 撤销前面做出的选择，恢复递归调用节点状态，以尝试其他选择
			path = path[:len(path)-1]
		}
	}

	backtrack(0, nums, make([]int, 0, 3))
	return result
}

// 这个实现版本，在循环中不再动态过滤出当前阶段对应的选择列表，
// 而是在发起下一个阶段的调用前后，更改和恢复选择列表
func permute3(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func(int, []int, []int)
	// 不能像下面这样声明与定义一行语句搞定，因为函数内部会递归调用其自身，
	// 导致报告：undefined: backtrack。必须先声明，再定义。
	// backtrack := func(k int, options []int, path []int) {
	backtrack = func(k int, options []int, path []int) {
		if k == len(nums) {
			// path 虽然是值传递（slice header 按值复制），但底层数组是共享的：
			// 初始 cap=3，整棵递归树上所有分支的 append 都复用同一块数组，
			// 后续分支会继续改写它。所以这里必须拷贝一份，再存入 result。
			result = append(result, append([]int(nil), path...))
			return
		}

		for i, opt := range options {
			path = append(path, opt)
			// 只从 options 删除 i 位置的元素
			options = slices.Delete(options, i, i+1)
			backtrack(k+1, options, path)

			// 撤销前面做出的选择，恢复递归调用节点状态，以尝试其他选择
			path = path[:len(path)-1]

			// 这样追加实现的恢复，导致通过迭代推演第一个阶段的第二个选择时，
			// 所做的选择跟第一个选择相同，进而导致结果中出现重复的答案
			// options = append(options, opt)

			// Del:      options=[1,2,3], del index=1, options=[1,3]
			// restore:  options=[1,3], insert 2 at index=1, options=[1,2,3]
			options = slices.Insert(options, i, opt)
		}
	}

	backtrack(0, nums, make([]int, 0, 3))
	return result
}

// 迭代体 5 条语句：
// ① 做选择        path = append(path, opt)
// ② 缩窄选择列表  options = options[i+1:]
// ③ 递归调用      backtrack(k+1, options, path)   ← ②的结果作为实参，往下钻一层
// ④ 撤销选择      path = path[:len(path)-1]
// ⑤ 恢复选择列表  options = append([]int{opt}, options...)
//
// backtrack(k=0, options=[1,2,3], path=[])
// │
// │  迭代 i=0, opt=1:
// │    ① path=[]，append(1) → path=[1]
// │    ② options=options[1:]=[2,3]
// │    ③ backtrack(k=1, options=[2,3], path=[1])
// │    │   ← ③ 发起的新帧内容才往下一层缩进
// │    │  迭代 i=0, opt=2:
// │    │    ① path=[1]，append(2) → path=[1,2]
// │    │    ② options=options[1:]=[3]
// │    │    ③ backtrack(k=2, options=[3], path=[1,2])
// │    │    │
// │    │    │  迭代 i=0, opt=3:
// │    │    │    ① path=[1,2]，append(3) → path=[1,2,3]
// │    │    │    ② i+1 == len(options) → options=nil
// │    │    │    ③ backtrack(k=3, options=nil, path=[1,2,3])
// │    │    │    │  k==3 → result 记录 [1,2,3] ✓ 返回
// │    │    │    ④ path=[1,2]
// │    │    │    ⑤ options=append([3],nil)=[3]  ✓
// │    │    │    range 结束 → 返回
// │    │    ④ path=[1]
// │    │    ⑤ options=append([2],[3])=[2,3]  ✓
// │    │
// │    │  迭代 i=1, opt=3:
// │    │    ① path=[1]，append(3) → path=[1,3]
// │    │    ② i+1 == len(options) → options=nil   ✗ 错点1：应为 [2]（删 opt=3 本身还剩 [2]），nil 导致该分支空手而归
// │    │    ③ backtrack(k=2, options=nil, path=[1,3])
// │    │    │  k=2≠3，range nil 无循环 → 空手而归
// │    │    ④ path=[1]
// │    │    ⑤ options=append([3],nil)=[3]   ✗ 错点2：恢复不完整（应为 [2,3]），元素 2 丢了；无后果：i=1 已是最后一轮，循环结束即返回，父帧不受影响
// │    │    range 结束 → 返回
// │    ④ path=[]
// │    ⑤ options=append([1],[2,3])=[1,2,3]  ✓ 恢复完整
// │
// │  迭代 i=1, opt=2:
// │    ① path=[]，append(2) → path=[2]
// │    ② options=options[2:]=[3]      ✗ 错点3：丢掉前缀 [1]（应删 i 本身，剩 [1,3]）
// │    ③ backtrack(k=1, options=[3], path=[2])
// │    │
// │    │  迭代 i=0, opt=3:
// │    │    ① path=[2]，append(3) → path=[2,3]
// │    │    ② i+1 == len(options) → options=nil
// │    │    ③ backtrack(k=2, options=nil, path=[2,3])
// │    │    │  k=2≠3 → 空手而归
// │    │    ④ path=[2]
// │    │    ⑤ options=append([3],nil)=[3]
// │    │    range 结束 → 返回
// │    ④ path=[]
// │    ⑤ options=append([2],[3])=[2,3]   ✗ 错点4：元素 1 永久丢失（应为 [1,2,3]）→ 直接导致下一轮 i=2 越界
// │
// │  迭代 i=2, opt=3:
// │    ① path=[]，append(3) → path=[3]
// │    ② i+1=3，len(options)=2（此时 options=[2,3]），3≠2  ✗ 错点5：原本我认为要么i+1<len，要么i+1==len，不可能大于
// │       → options=options[3:] → 💥 panic: slice bounds out of range [3:2]
//
// 疑问：i=1 这个迭代第 ⑤ 步已把 options 改成 [2,3]，为什么 i=2 这轮执行还能访问索引 2？
// Go 规范：range 表达式只求值一次。求值得到的是 slice 的 header（指向底层数组的指针 + len），
// 这就是快照。它存在编译器生成的隐藏临时变量里，源码里没有名字、也访问不到，就像按值传参的
// 形参。编译器实际生成的大致是这样：
//
//   _snapshot := options              // 隐藏：进入循环时把 header 拷走
//   for _i := 0; _i < len(_snapshot); _i++ {
//       i := _i; opt := _snapshot[_i] // 循环变量每轮重赋值，元素从快照现读
//       ...循环体原样展开，其中的 options 仍是你那个变量...
//   }
//
// 所以运行时有「两套 options」：迭代计划只看 _snapshot（len=3 定死 i=0,1,2 三轮），
// 循环体内的 ②⑤ 只改「变量 options」。i=2 轮照常到来，但 ② 读的是 options 变量
// 当前值 [2,3]：len=2，options[3:] 起点 3 超长 → 越界：
//
// 			if i+1 == len(options) {
//				options = nil
//			} else {
//				options = options[i+1:]
//			}
//
// 区分：原地改元素（如 options[0]=9，共享底层数组）后续轮次可见；给变量换 header 循环看不见。

// 这个实现版本有严重且难以诊断的 bug。主要 bug 如下（从上面注释中的完整调用树提取出来的）：
// backtrack(k=0, options=[1,2,3], path=[])
// │
// │  迭代 i=0, opt=1:
// │    ① path=[]，append(1) → path=[1]
// │    ② options=options[1:]=[2,3]
// │    ③ backtrack(k=1, options=[2,3], path=[1])
// │    ④ path=[]
// │    ⑤ options=append([1],[2,3])=[1,2,3]  ✓ 恢复完整
// │
// │  迭代 i=1, opt=2:
// │    ① path=[]，append(2) → path=[2]
// │    ② options=options[2:]=[3]      ✗ 错点3：丢掉前缀 [1]（应删 i 本身，剩 [1,3]）
// │    ③ backtrack(k=1, options=[3], path=[2])
// │    ④ path=[]
// │    ⑤ options=append([2],[3])=[2,3]   ✗ 错点4：元素 1 永久丢失（应为 [1,2,3]）→ 直接导致下一轮 i=2 越界
// │
// │  迭代 i=2, opt=3:
// │    ① path=[]，append(3) → path=[3]
// │    ② i+1=3，len(options)=2（此时 options=[2,3]），3≠2
// │       → options=options[3:] → 💥 panic: slice bounds out of range [3:2]
//
// 导致上述 bug 的思考路径：思考这个实现思路时，脑袋里想的场景是第一次循环 i = 1，
// 在这次循环中，确实正确恢复了 options，就以为这个思路是可行的。但当第二次循环
// i = 2 时，循环变量 opt 只可能保存当前循环访问的元素，不可能保存前一次循环访问的元素，
// 因此恢复时就不可能恢复前一/几次循环访问的元素。
//
// 这个实现版本中不止这一个 bug，还有其他 bug，具体请看前面注释中的完整调用树。
func permute_serious_hardtodiagnosebug(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func(int, []int, []int)
	// 不能像下面这样声明与定义一行语句搞定，因为函数内部会递归调用其自身，
	// 导致报告：undefined: backtrack。必须先声明，再定义。
	// backtrack := func(k int, options []int, path []int) {
	backtrack = func(k int, options []int, path []int) {
		if k == len(nums) {
			// path 虽然是值传递（slice header 按值复制），但底层数组是共享的：
			// 初始 cap=3，整棵递归树上所有分支的 append 都复用同一块数组，
			// 后续分支会继续改写它。所以这里必须拷贝一份，再存入 result。
			result = append(result, append([]int(nil), path...))
			return
		}

		for i, opt := range options {
			path = append(path, opt)

			if i+1 == len(options) {
				options = nil
			} else {
				options = options[i+1:]
			}

			backtrack(k+1, options, path)

			// 撤销前面做出的选择，恢复递归调用节点状态，以尝试其他选择
			path = path[:len(path)-1]
			options = append([]int{opt}, options...)
		}
	}

	backtrack(0, nums, make([]int, 0, 3))
	return result
}

func TestPermute(t *testing.T) {
	got := permute([]int{1, 2, 3})
	want := [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
		{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	}
	assert.ElementsMatch(t, want, got)
}
