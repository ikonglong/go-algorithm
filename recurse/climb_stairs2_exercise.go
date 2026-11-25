package recurse

// https://leetcode.cn/problems/three-steps-problem-lcci/
// 楼梯有 n 阶台阶，小孩一次可以上 1 阶、2 阶或 3 阶

// 递归实现。递推公式：f(n) = f(n-1) + f(n-2) + f(n-3), n>=3
// f(0)=1, f(1)=1, f(2)=2
func waysToStep(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return (waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)) % 1000000007
}
