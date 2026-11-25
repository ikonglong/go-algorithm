package dp

// https://leetcode.cn/problems/climbing-stairs/
// 逆向思考，假设现在已经爬到了楼顶，回看自己最后一步是怎么走的，只有两种走法：爬一阶，爬两阶

// 有多少级台阶，就有多少个处理阶段？
// 不对，这里无法先确定有多少个阶段，因为步长不确定。应该说，走了多少步就有多少个阶段
//
// 状态表：dp[n+1]，n 是台阶总数，dp[i] 是走完 i 个台阶有多少走法
// 状态转移方程：dp[i] = dp[i-1] + dp[i-2]，i-2>=0
func climbStairs_(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	dp := make([]int, n+1)

	// 初始化
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
