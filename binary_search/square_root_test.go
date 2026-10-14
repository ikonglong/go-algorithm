package binary_search

// https://leetcode.cn/problems/sqrtx/
// 解空间为：[1, x]，可进一步优化为 [1, x/2+1]，
// 然后二分查找其平方小于等于 X 的最后一个数

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	lo, hi := 1, x/2+1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// 先判断 mid 是否属于解空间
		midSquare := mid * mid
		if midSquare <= x {
			// 再判断是否是真解
			// 想要 mid+1<=hi => mid<=hi-1 => 想到要处理 mid == hi 边界情形
			if mid == hi || (mid+1)*(mid+1) > x {
				return mid
			} else {
				lo = mid + 1
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}
