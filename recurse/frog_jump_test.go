package recurse

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// 平台有 num 个小格子，每次可以选择跳 一个格子 或者 两个格子。

// 1. 跳 n 个格子跟跳 n-1 个格子的解决思路一样
// 2. n 个格子的所有走法，等于第一步走1格、剩余 n-1 格的所有走法，加上第一步走2格、剩余 n-2 格的所有走法
// 递推公式：f(n) = f(n-1) + f(n-2)
// 递归终止条件：要求 n-2>=0 => n>=2，f(0)=1, f(1)=1
// 只有0个格子时，站在原地不动或走0格也是一种走法
//
// 这个实现版本执行超时。还可以采用动态规划的逆向推导解法
func trainWays(num int) int {
	if num < 2 {
		return 1
	}
	return trainWays(num-1) + trainWays(num-2)
}

// 动态规划思路:
// 1. 每个台阶对应一个阶段。有 n 个台阶就有 n 个阶段
// 2. 反/逆向推导：走到第 n 个台阶，最后一步有 2 个选择，分别走 1、2 阶
//
// 状态转移表和方程：
// dp[n+1], dp[i] 表示走到第 i 阶的走法总数
// dp[i] = dp[i-1] + dp[i-2]
func trainWays2(num int) int {
	// 因为递推公式中访问的最小下标为 n-2，所以 n>=2 才可直接套用递推公式。
	// 因此选择直接初始化 dp[0]、dp[1]
	if num < 2 {
		return 1
	}
	dp := make([]int, num+1)
	dp[0] = 1
	dp[1] = 1
	mod := 1000000007
	for i := 2; i <= num; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[num]
}
