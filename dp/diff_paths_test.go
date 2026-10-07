package dp

// https://leetcode.cn/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
	// 创建状态转移表，每个单元格存储的是到达这个单元格的路径数量
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 因为到达起点的路径只有一条，所以初始化为 1。
	// 第一行没有上方相邻单元格，只有左侧相邻单元格，只能初始化为 1；
	// 第一列没有左侧相邻单元格，只有上方相邻单元格，也只能初始化为 1
	// 初始化第一行
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 初始化第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
