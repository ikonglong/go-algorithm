package slidingwin

import "math"

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
//
// 示例：'bacabccb', 'abcd', 'kkk'

// 算法思路：把滑动窗口想象成一支队伍，字符从右侧入口进入，从左侧出口离开。
// p 指向窗口内下一个可能离开的字符；q 指向右侧入口外、下一个等待进入的候选字符。
// 因此窗口的下标范围是左闭右开的 [p, q)，q 不属于窗口，
// 窗口内最后一个字符是 arr[q-1]。
//
// 每轮循环开始和结束时都保持：[p, q) 中没有重复字符，
// charSet 恰好记录窗口中的字符。
// 1. 初始时 p=q=0，窗口范围 [0, 0) 为空。
// 2. 如果候选字符 arr[q] 已在窗口中，就让它留在入口处，并不断从左侧移出字符，
//    直到窗口中不再有相同字符。
// 3. 候选字符不再冲突后，将它记录到 charSet，并右移 q，使它进入窗口；
//    此时窗口重新满足不变量，再用 q-p 更新最大长度。
//    收缩窗口只会使窗口变短，因此无须在收缩时更新最大长度。

// 每轮处理一个候选字符：先消除冲突，再让它进入窗口。
func dismantlingAction(arr string) int {
	if arr == "" {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	p, q := 0, 0
	maxLen := 0
	charSet := make(map[byte]struct{})
	for q < len(arr) {
		c := arr[q]

		// 候选字符留在入口处；如果重复，就从出口收缩窗口，直到它可以进入。
		for {
			if _, found := charSet[c]; !found {
				break
			}
			delete(charSet, arr[p])
			p++
		}

		charSet[c] = struct{}{}
		q++
		maxLen = max(maxLen, q-p)
	}
	return maxLen
}

// 破除“循环内每一个分支必须推进循环变量”这个执念后的最简版本
func dismantlingAction2(arr string) int {
	if arr == "" {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	p, q := 0, 0
	maxLen := 0
	charSet := make(map[byte]struct{})
	for q < len(arr) {
		c := arr[q]
		if _, found := charSet[c]; !found {
			charSet[c] = struct{}{}
			q++
			maxLen = max(maxLen, q-p)
			continue
		}

		// 如果最后一次迭代走的是「有重复」这个分支，那么循环结束后是否还需要再计算一次 maxLen？
		// 不需要。【关键在于：发现重复时，q 不会向右移动。】因此，即使最后一个字符第一次检查时
		// 发现重复，这也不是循环真正的“最后一次执行”。程序会移动 p，然后再次检查同一个 arr[q]。
		for found := true; found; _, found = charSet[c] {
			delete(charSet, arr[p])
			p++
		}
	}
	return maxLen
}

// 还算简洁的实现版本，但循环中的两个分支还有一些重复逻辑。
// 重复的根源在于，我试图在每个分支中都移动 q。但其实如果重复，
// 只将 P 移动至子串中最后一个重复字符之后，在下一次迭代中就会移动 Q，
// 虽然多了一次迭代，但代码更简洁，更不容易埋 bug。
// 这也挖出了我脑子中的错误认知：只要循环中的某个分支没有推进循环变量，就一定有问题。
func dismantlingAction3(arr string) int {
	if arr == "" {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	p, q := 0, 0
	maxLen := 0
	charSet := make(map[byte]struct{})
	for q < len(arr) {
		// 待处理部分的第一个字符跟当前不重复子串中的字符是否重复？
		if _, found := charSet[arr[q]]; found {
			// 结束当前子串
			maxLen = max(q-p, maxLen)

			// 移动 p 到子串中最后一个重复字符的下一个字符，
			// 开始新的子串，p 依然指向子串头
			for ; found; _, found = charSet[arr[q]] {
				delete(charSet, arr[p])
				p++
			}
			charSet[arr[q]] = struct{}{}
			q++
		} else {
			charSet[arr[q]] = struct{}{}
			q++
		}
	}
	maxLen = max(maxLen, q-p)
	return maxLen
}

// 算法思路：
// 概述：用两个指针，分别指向不重复连续子串、未处理部分的第一个元素
// 1. 初始时，假定第一个字符就是不重复连续子串。
//    左指针指向第一个字符；右指针指向第二个字符，即待处理部分的第一个字符
// 2. 逐个遍历待处理部分的每个字符
// 	    如果当前字符在不重复子串中存在，则需移动左指针到重复字符后一个字符

// 这个实现版本尚可接受。与后一个实现版本相比，移除了不必要的、且会带来麻烦的额外 count
func dismantlingAction4(arr string) int {
	if arr == "" {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	maxLen := math.MinInt
	chars := make(map[byte]struct{})
	chars[arr[0]] = struct{}{}
	p, q := 0, 1
	for q < len(arr) {
		if _, found := chars[arr[q]]; found {
			// 发现重复，先结束当前不重复子串
			maxLen = max(q-p, maxLen)

			// 移动 p 到重复字符的下一个字符，开始一个新的不重复子串，
			for ; found; _, found = chars[arr[q]] {
				delete(chars, arr[p])
				p++
			}

			// 始终准确保存窗口 arr[p:q] 中的字符。很容易忘
			chars[arr[q]] = struct{}{}
			q++
		} else {
			chars[arr[q]] = struct{}{}
			q++
		}
	}
	maxLen = max(q-p, maxLen)
	return maxLen
}

func dismantlingAction4_notGood(arr string) int {
	if arr == "" {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}

	maxLen := math.MinInt
	// 观察代码会发现，不需要 count，因为 q-p 就是子串的长度。
	// 额外定义 count，就需要在前一个子串结束、后一个子串开始时
	// 先结算 count，再重新初始化 count，这里很容易埋 bug
	count := 1
	chars := make(map[byte]struct{})
	chars[arr[0]] = struct{}{}
	p, q := 0, 1
	for q < len(arr) {
		if _, found := chars[arr[q]]; found {
			// 发现重复，先结束当前不重复子串
			count = q - p
			maxLen = max(count, maxLen)

			// 移动 p 到重复字符的下一个字符，开始一个新的不重复子串
			for ; found; _, found = chars[arr[q]] {
				delete(chars, arr[p])
				p++
			}

			// 这行代码很别扭，且埋了难以排查的 bug。
			// 别扭的原因是，对于输入 'kkk'，走到这里 p=q=1，滑窗 [1,1) 是个空窗，
			// 并不包含 1。但当前分支处理的是 arr[q]，且 p 和 q 都为 1，业务逻辑上
			// 也应该包含。为了包含就针对具体情形硬凑，搞出 count=q-p+1，写这样代码时
			// 就感觉别扭，因为在这里 q 的含义变了，不再指向未处理部分的第一个，而指向
			// 当前子串的最后一个。这就改变了 q 的定义，埋下了 bug。
			//
			// 正确的思考路径是：先移动 q，再计算 count=q-p。这样 q 的定义与使用
			// 就全局一致了
			//count = q - p + 1

			// 始终准确保存窗口 arr[p:q] 中的字符。很容易忘
			chars[arr[q]] = struct{}{}
			q++
			count = q - p
			// 不能这么写，因为并非所有情形循环结束后都是p=q=x
			//chars[arr[p]] = struct{}{}
		} else {
			chars[arr[q]] = struct{}{}
			count++
			q++
		}
	}
	maxLen = max(count, maxLen)
	return maxLen
}
