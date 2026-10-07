package dp

import "math"

// https://leetcode.cn/problems/triangle/description/

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	// 初始化第一行，或说三角形的顶角
	//dp[0][0] = 1 // bug
	dp[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		n := len(dp[i])
		for j := 0; j < n; j++ {
			if j == 0 { // 当前行 0 元素只对应上一行 0 元素
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == n-1 { // 当前行比上一行长1，因此最后元素只对应j-1元素
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	minSum := math.MaxInt
	for _, v := range dp[m-1] {
		minSum = min(minSum, v)
	}
	return minSum
}
