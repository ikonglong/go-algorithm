package dp

import "math"

// https://leetcode.cn/problems/coin-change/description/
// 既然是硬币面额，自然没有负数

// 算法思路：
// - 多阶段决策模型
// 1. 每个阶段处理一个硬币面额
// 2. 每个阶段决定一个硬币面额是否参与凑数
// 3. 路径就是状态转移表
// - 状态转移方程
// dp[i][j], 0<=j<=sum(coins[i]), 单元格是凑成某个总金额所需的最少硬币数
// dp[i][j] = min(dp[i-1][j],
// dp[i-1][j-coins[i]]+1, dp[i-1][j-2*coins[i]]+2, ..., dp[i-1][j-k*coins[i]]+k)
// k = j/coins[i]

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount <= 0 {
		return 0
	}

	n := len(coins)
	dp := make([][]int, n)
	for i := range dp {
		// 这样写不好。把 amount 看作背包问题中背包的最大承重才好
		// 第二维长度不能是 amount+1，因为第一个面值可能大于 amount
		// dp[i] = make([]int, max(amount, coins[0])+1)
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	// 不能再像 0-1 背包问题那样初始化第一行了
	// dp[0][0] = 0 // 第一个不参与
	// // 这样写不好。把 amount 看作背包问题中背包的最大承重才好
	// // dp[0][coins[0]] = 1 // 参与
	// if coins[0] <= amount {
	// 	dp[0][coins[0]] = 1
	// }
	for k := 0; k <= amount/coins[0]; k++ {
		dp[0][k*coins[0]] = k
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= amount; j++ {
			for k := 0; k <= j/coins[i]; k++ {
				// if j-k*coins[i] >= 0 {
				// 	dp[i][j] = min(dp[i][j], dp[i-1][j-k*coins[i]]+k)
				// }
				if dp[i-1][j-k*coins[i]] != math.MaxInt {
					dp[i][j] = min(dp[i][j], dp[i-1][j-k*coins[i]]+k)
				}
			}

			// 不管走哪个分支，只要父节点是 dp[i-1][j]，就表示 coins[i] 不参与
			// if j < coins[i] {
			// 	dp[i][j] = dp[i-1][j]
			// } else {
			// 	dp[i][j] = min(dp[i-1][j], dp[i-1][j-coins[i]]+1)
			// }
		}
	}

	if dp[n-1][amount] == math.MaxInt {
		return -1
	}
	return dp[n-1][amount]
}

// 使用爬楼梯模型实现零钱兑换
func coinChange2(coins []int, amount int) int {
	k := len(coins) // 有几种面值
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0 // 凑成 0 金额至少需要 0 个硬币
	for i := 1; i <= amount; i++ {
		// 将状态每推进一个单位，都有 k 个选择，每个选择都需要判断处理
		// 有 k 种面值，就有 k 个选择
		for j := 0; j < k; j++ {
			// 反向推导，当前状态为 i，前一个状态就为 i-coins[j]
			if i-coins[j] >= 0 &&
				dp[i-coins[j]] != math.MaxInt && // 前一状态是否可达
				dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1 // 不要 +=
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
