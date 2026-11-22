package twopointer

// 乍一想我以为应该采用类似反转链表的方法，需要 previous、current 指针，仔细分析后发现并不可行。
// 算法思路：
// 创建两个指针 low、high，分别指向第一个、最后一个字符。
// 使用这两个指针同步遍历数组，直到 low >= high
// 1. 交换对应位置的元素
// 2. low 向右移动，high 向左移动

func reverseString(s []byte) {
	// 这个检查可选
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
