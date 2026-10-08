package recurse

// https://leetcode.cn/problems/three-steps-problem-lcci/
// 楼梯有 n 阶台阶，小孩一次可以上 1 阶、2 阶或 3 阶
//
// 1. 爬 n 个台阶跟上 n-1 个台阶的解决思路一样
// 2. 爬 n 级台阶的爬法 等于
// 		第一步爬 1 阶、剩余 n-1 阶的爬法，加上
//      第二步爬 2 阶、剩余 n-2 阶的爬法，加上
//      第三步爬 3 阶、剩余 n-3 阶的爬法
//    n-3>=0 => n>=3
//    f(0) = 1
//    f(1) = 1
//    f(2) = 2
//    f(n) = f(n-1)+f(n-2)+f(n-3)

// 这个实现版本执行超时
func waysToStep(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	}
	return waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)
}

// 动态规划思路:
// 1. 每个台阶对应一个阶段。有 n 个台阶就有 n 个阶段
// 2. 反/逆向推导：走到第 n 个台阶，最后一步有 3 个选择，分别走 1、2、3 阶
//
// 状态转移表和方程：
// dp[n+1], dp[i] 表示走到第 i 阶的走法总数
// dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
func waysToStep2(n int) int {
	mod := 1000000007
	dp := make([]int, n+1) // 默认值 0 表示不可达
	dp[0] = 1              // 走到第 0 阶，走 0 步或保持原地不动，这是一种走法
	for i := 1; i <= n; i++ {
		if i-1 >= 0 {
			// 即使 dp[i-1] 不可达，也可不判断，直接赋值，因为其默认值为 0，赋值了结果也符合预期
			dp[i] = (dp[i] + dp[i-1]) % mod
		}
		if i-2 >= 0 {
			dp[i] = (dp[i] + dp[i-2]) % mod
		}
		if i-3 >= 0 {
			dp[i] = (dp[i] + dp[i-3]) % mod
		}
	}
	return dp[n]
}

// 因为递推公式中访问的最小下标为 n-3，所以 n>=3 才可直接套用递推公式。
// 因此选择直接初始化 dp[0]、dp[1]、dp[2]
func waysToStep3(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	}
	mod := 1000000007
	dp := make([]int, n+1) // 默认值 0 表示不可达
	dp[0] = 1              // 走到第 0 阶，走 0 步或保持原地不动，这是一种走法
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % mod
	}
	return dp[n]
}
