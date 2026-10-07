package dp

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
// 分析题目，建立多阶段决策模型：
// 1. 有多少天就有多少个决策阶段，每个决策阶段决定当天是买入、卖出还是什么也不做。
//    同一天不能买入卖出。买入、卖出数量都只能是一股
// 2. 如果手上持有股票，则在卖出之前不能再购买其他股票。这就意味着任何时刻，要么未持有任何股票，要么仅持有一只股票。
// 3. 此外，阶段间的相互影响是：持有股票才可卖出，不持有才可买入
//
// 状态转移方程：
// dp[n][2] 状态转移表记录每个阶段即每天的状态。每天有两种状态：持有、不持有股票
// dp[i][0] 表示第 i 天不持有股票赚到的最大利润
// dp[i][1] 表示第 i 天持有股票赚到的最大利润
//
// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
// dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 初始化第一天
	dp[0][0] = 0          // 第一天不持有股票，意味着不操作
	dp[0][1] = -prices[0] // 第一天持有股票，则意味着买入
	for i := 1; i < n; i++ {
		// 当天结束时不持有
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		// 当天结束时持有
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return max(dp[n-1][0], dp[n-1][1])
}
