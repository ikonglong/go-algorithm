package recurse

// https://leetcode.cn/problems/recursive-mulitply-lcci/
// 1. a×n 跟 a×(n-1) 解决思路一样。
// 2. 递推公式：a×n = a + a×(n-1)
//    终止条件：n-1>=0 => n>=1, a×0=0

func multiply(A int, B int) int {
	return recurse(A, B)
}

func recurse(A int, times int) int {
	if times == 0 {
		return 0
	}
	return A + recurse(A, times-1)
}
