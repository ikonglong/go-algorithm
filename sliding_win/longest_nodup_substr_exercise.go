package slidingwin

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
//
// 示例：'bacabccb', 'abcd', 'kkk'

func dismantlingAction_(arr string) int {
	//maxLen := math.MinInt
	maxLen := 0
	p, q := 0, 0 // [p, q) 空区间
	charSet := make(map[byte]struct{})

	for q < len(arr) {
		c := arr[q]

		// 注意这里不是直接删除 c，而是循环删除 arr[p]，直到 c 被移出窗口左端
		for _, found := charSet[c]; found; _, found = charSet[c] {
			delete(charSet, arr[p])
			p++
		}

		charSet[c] = struct{}{}
		q++
		maxLen = max(maxLen, q-p)
	}

	return maxLen
}
