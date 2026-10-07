package dp

import "math"

// https://leetcode.cn/problems/jian-sheng-zi-lcof/description/

// 类比爬楼梯问题：竹子长度就是台阶总数，每次走几步就是多长砍一刀。
// 题目没有明确多长砍一刀，可以认为每一个可能的长度都可砍一刀。
// 也就是说，可砍一刀的长度范围为 [1, bamboo_len]
//
// 状态转移表和方程：
// dp[bamboo_len+1]，长度在[0, bamboo_len]范围内的竹子切分后的最大乘积
// dp[i]，长度为 i 的竹子切分后所有段的最大乘积
// dp[i] = max(dp[i-1]*1, dp[i-2]*2, ..., dp[i-i]*i)
//
// dp[0] = 1，为什么？
// dp[0] 不是一个真实的子问题：没有长度 0 的竹子可切。它只是当竹子被全部切完、一点都不剩时，
// 代表「剩余部分 」（此时为空）对乘积的贡献：既然是「没有段可乘，或乘以空」，就不该改变已有
// 的乘积，所以其值为乘法单位元 1。
// 三个你早就接受的类比：
// 	- 0! = 1
// 	- a⁰ = 1
// 	- 一个空列表的连乘 = 1
// 它们的共同点不是「算出来等于 1」，而是「什么都没乘」时，不能改变已有结果。乘法里唯一有这个
// 性质的数就是 1（加法里是 0）。所以乘积型递推的空基座天然是 1。
//
// 为什么很多人直觉上写成 dp[0] = 1？
// 因为加法型 DP 的空基座是 0（dp[0]=0，加 0 不变），比如爬楼梯/凑硬币。
// 但这里子问题的值存的是乘积不是和，空基座就得跟着运算换成 1。运算变，单位元就变。
// 注意，但并非所有加法型 DP 的空基座一定初始化为 0，请看 crack_num_test.go

func cuttingBamboo(bamboo_len int) int {
	// 顶层问题或叫原问题要求「至少切一刀」，但若 2、3 是原问题分解出来的子问题，
	// 或叫 dp 中的子问题/子段，则可不切，因为切了子段可能就得不到最优解
	if bamboo_len == 1 {
		return 1
	}
	if bamboo_len == 2 {
		return 1 // 1*1
	}
	if bamboo_len == 3 {
		return 2 // 1*2
	}
	dp := make([]int, bamboo_len+1)
	for i := range dp {
		dp[i] = math.MinInt
	}

	// dp[i] 表示「长为 i 的竹子（作为子段，可不切）的最大乘积」

	dp[0] = 1
	for i := 1; i <= bamboo_len; i++ {
		// 对于子段 i，可以不切，切了可能得不到最优解，因此 j 可取到 i
		for j := 1; j <= i; j++ {
			// if dp[i] < dp[i-j]*j {
			// 	dp[i] = dp[i-j] * j
			// }
			dp[i] = max(dp[i], dp[i-j]*j)
		}
	}
	return dp[bamboo_len]
}

func cuttingBamboo2(bamboo_len int) int {
	// 顶层问题或叫原问题要求「至少切一刀」，但若 2、3 是原问题分解出来的子问题，
	// 或叫 dp 中的子问题/子段，则可不切，因为切了子段可能就得不到最优解
	if bamboo_len == 1 {
		return 1
	}
	if bamboo_len == 2 {
		return 1 // 1*1
	}
	if bamboo_len == 3 {
		return 2 // 1*2
	}
	dp := make([]int, bamboo_len+1)
	for i := range dp {
		dp[i] = math.MinInt
	}

	// dp[i] 表示「长为 i 的竹子（作为子段，可不切）的最大乘积」
	// dp[0] 不是一个真实的子问题：没有长度 0 的竹子可切。它只是当竹子被全部切完、一点都不剩时，
	// 代表「剩余部分 」（此时为空）对乘积的贡献：既然是「没有段可乘，或乘以空」，就不该改变已有
	// 的乘积，所以其值为乘法单位元 1。
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2 // 关键：子段 2 可不切分
	dp[3] = 3 // 关键：子段 3 可不切分
	for i := 1; i <= bamboo_len; i++ {
		// 对于子段 i，可以不切，切了可能得不到最优解，因此 j 可取到 i
		for j := 1; j <= i/2; j++ { // j 到 i/2 即可，避免重复
			if v := dp[i-j] * j; v > dp[i] {
				dp[i] = v
			}
		}
	}
	return dp[bamboo_len]
}
