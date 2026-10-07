package dp

import "math"

func jewelleryValue(frame [][]int) int {
	m := len(frame)
	n := len(frame[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = frame[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + frame[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + frame[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + frame[i][j]
		}
	}

	maxVal := math.MinInt
	for _, v := range dp[m-1] {
		maxVal = max(maxVal, v)
	}
	return maxVal
}
