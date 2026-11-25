package dp

// https://leetcode.cn/problems/climbing-stairs/
// 每次只能爬 1 或 2 个台阶
//
// 对比递归实现 recurse/climb_stairs.go

// 有多少级台阶，就有多少个处理阶段？
// 不对，这里无法先确定有多少个阶段，因为步长不确定。应该说，走了多少步就有多少个阶段
//
// 状态表：dp[n+1]，n 是台阶总数，dp[i] 是走完 i 个台阶有几种走法
// 状态转移方程：
// dp[i] = dp[i-1] + dp[i-2]，i-2>=0
//
// dp[2] = dp[1] + dp[0] 用问题领域语言解释这个式子：
// 到达第 2 阶的所有走法，可以按最后一步跨了几阶，分成两类：
// 最后一步跨 1 阶、最后一步跨 2 阶
// 注意，这里加的是两类走法的数量，不是把“走到第 1 阶”和“走到第 0 阶”两个过程接起来。

func climbStairs(n int) int {
	dp := make([]int, n+1)
	// 走到第 0 阶就是站在原地不动，即走 0 步或对应的步长序列为空，
	// 但这就是走到第 0 阶的一个方法，因此走到第 0 阶的方法为 1，而不是 0。
	// 走到第 0 阶只需走 0 步，跟走到第 0 阶的方法数是 1，并不矛盾。
	dp[0] = 1
	for i := 1; i <= n; i++ {
		// 原本要做的计算是：dp[i] = dp[i-1] + dp[i-2]
		// 但只有 dp[i-1] 和 dp[i-2] 存在或合法时，才能完整计算。
		// 只有 dp[i-1] 存在时就只算它一项
		if i-1 >= 0 {
			dp[i] += dp[i-1]
		}
		if i-2 >= 0 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func climbStairs2(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
