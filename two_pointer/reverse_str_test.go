package twopointer

// 算法思路：
// 乍一想我以为应该采用类似反转链表的方法，需要 previous、current 指针，仔细分析后发现并不可行。
// 正确的方法是创建两个指针 left 和 right，分别指向第一个字符和最后一个字符。
// 1. 交换 s[left]、s[right]
// 2. left 向右移动，right 向左移动

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
