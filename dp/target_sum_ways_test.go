package dp

import "math"

// https://leetcode.cn/problems/target-sum/

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	n := len(nums)
	// 为正数状态创建状态转移表。
	// 列为表达式可达的结果，单元格为到达这个结果的路径个数
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 1000+1)
		// for j := range dp[i] {
		// 	dp[i][j] = math.MaxInt
		// }
	}
	// 为负数状态创建状态转移表
	dp2 := make([][]int, n)
	for i := range n {
		dp2[i] = make([]int, 1000+1)
		// for j := range dp2[i] {
		// 	dp2[i][j] = math.MaxInt
		// }
	}

	// 初始化第一行
	if nums[0] == 0 {
		dp[0][nums[0]] = 2
	} else {
		dp[0][nums[0]] = 1  // +
		dp2[0][nums[0]] = 1 // -
	}

	// 状态转移方程：
	// dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
	for i := 1; i < n; i++ {
		for j := -1000; j <= 1000; j++ {
			// if j-nums[i] < -1000 || j+nums[i] > 1000 {
			// 	continue
			// }
			v1, v2 := 0, 0
			if math.Abs(float64(j-nums[i])) <= 1000 {
				if j-nums[i] >= 0 {
					v1 = dp[i-1][j-nums[i]]
				} else {
					v1 = dp2[i-1][-(j - nums[i])]
				}
			}
			if math.Abs(float64(j+nums[i])) <= 1000 {
				if j+nums[i] >= 0 {
					v2 = dp[i-1][j+nums[i]]
				} else {
					v2 = dp2[i-1][-(j + nums[i])]
				}
			}

			if j >= 0 {
				dp[i][j] = v1 + v2
			} else {
				dp2[i][-j] = v1 + v2
			}
		}
	}

	if target >= 0 {
		return dp[n-1][target]
	}
	return dp2[n-1][-target]
}
