package twopointer

// https://leetcode.cn/problems/reverse-string/

func reverseString_(s []byte) {
	// 这个检查可不写
	if len(s) == 0 {
		return
	}

	lo, hi := 0, len(s)-1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
