package slidingwin

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
//
// 示例：'bacabccb', 'abcd', 'kkk'

// 算法思路：
// 把滑动窗口想象成一条走廊，字符只能从右侧进入，从左侧离开。
// p 是窗口左闭边界；q 是窗口的右开边界，指向待处理的第一个字符，即右侧入口外下一个等待进入的候选字符。
//
// 用一个 map charSet 记录窗口中的字符。
// 初始时 p=q=0，窗口 [0, 0) 为空。
// 使用 q 遍历整个字符数组。每次迭代只处理一个候选字符：先消除重复，再让它进入窗口。每次迭代开始和结束时都保持 [p, q) 中没有重复字符
// 1. 如果候选字符 arr[q] 跟窗口中的重复，q 先保持不动。
//    循环右移窗口左端，移出左端字符，直到相同字符移出窗口；
//    收缩窗口只会使窗口变短，因此无须在收缩时更新最大长度。
// 2. 候选字符不再重复后，将它记录到 charSet，并右移 q，使它进入窗口；
//    此时窗口重新满足不变量，再用 q-p 更新最大长度。

func dismantlingAction(arr string) int {
	// Bug: arr 是空串时不对
	// maxLen := math.MinInt
	maxLen := 0
	p, q := 0, 0 // [p, q) 空区间
	charSet := make(map[byte]struct{})

	for q < len(arr) {
		c := arr[q]

		// 注意这里不是直接删除 c，而是循环删除 arr[p]，直到 c 被移出窗口左端
		// d, b, a, c, a, ...
		// p           q
		// arr[q] 待进入窗口，但跟窗口中的字符重复，
		// 需要循环移动 p，直到将窗口中的相同字符 a 移出窗口
		// Bug: 该死的，循环变量 found 没有更新，看到了吗？
		for _, found := charSet[c]; found; {
			delete(charSet, arr[p])
			p++
			// 要么在 for 语句中更新循环变量，要么在这里更新
			_, found = charSet[c]
		}

		charSet[c] = struct{}{}
		q++
		maxLen = max(maxLen, q-p)
	}

	return maxLen
}
