package dp

// 把 amount 看作背包问题中的背包最大承重
// 组合数就是路径数
func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}

	n := len(coins)
	// 创建状态转移表
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		// for j := range dp[i] {
		// 	dp[i][j] = math.MinInt
		// }
	}
	// 初始化第一行
	for k := 0; k <= amount/coins[0]; k++ {
		dp[0][k*coins[0]] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= amount; j++ {
			for k := 0; k <= amount/coins[0]; k++ {
				if j-k*coins[i] >= 0 {
					dp[i][j] = dp[i][j] + dp[i-1][j-k*coins[i]]
				}
			}
		}
	}
	return dp[n-1][amount]
}
