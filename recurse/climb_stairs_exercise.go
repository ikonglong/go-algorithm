package recurse

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// 平台有 num 个小格子，每次可以选择跳 一个格子 或者 两个格子。

// 递归实现。递推公式：f(n) = f(n-1) + f(n-2), n>=2
func trainWays_(num int) int {
	if num < 2 {
		return 1
	}
	return trainWays_(num-1) + trainWays_(num-2)
}
