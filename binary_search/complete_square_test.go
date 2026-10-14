package binary_search

// https://leetcode.cn/problems/valid-perfect-square/
// 算法思路：
// 1. num 平方根的解空间 [1, num]，可进一步优化为 [1, x/2+1]
// 2. 在其中用二分查找找 num 的整数平方根

func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
