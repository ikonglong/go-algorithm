package dp

import "math"

// weights 表示物品列表，同时也表示每个物品的重量
func knapsackMaxVal(weights []int, values []int) int {
	n := len(weights)
	wCap := 9

	// 创建状态转移表格 dp
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, wCap+1)
		for j := range wCap + 1 {
			// 注意，这里不能用零值作为初始值，因为 value 0 是可达的，
			// 初始值也是零，就无法区分不可达的
			dp[i][j] = math.MinInt
		}
	}

	// 初始化第一行。对于物品 0，有两种选择：放，不放
	dp[0][0] = 0 // 不放
	if weights[0] <= wCap {
		dp[0][weights[0]] = values[0] // 放
	}

	// 逐行推导
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 上一行 w 是否到达了
			// 模拟从决策树的上一层 w 节点派生当前层节点
			if dp[i-1][w] >= 0 {
				// 两种选择：放、不放
				// 注意，这里必须比较当前单元格跟上一行单元格，取最大值。
				// 因为有可能其他决策树分支已经将当前单元格设置为更大的值了
				// 不放
				dp[i][w] = int(math.Max(float64(dp[i][w]), float64(dp[i-1][w])))
				// 注意，这里必须比较当前单元格跟上一行单元格，取最大值。
				// 因为有可能其他决策树分支已经将当前单元格设置为更大的值了
				// 放
				if w+weights[i] <= wCap { // 这个容易遗漏
					dp[i][w+weights[i]] = int(math.Max(float64(dp[i][w+weights[i]]), float64(dp[i-1][w]+values[i])))
				}
			}
		}
	}

	// 不可以。dp[i][w] 表示考虑了前 i+1 个物品，恰好达到重量 w 时能够获得的最大价值。
	// 列下标 w 按重量排序，但单元格中的价值不会随重量单调增加。重量越大，不代表价值越大。
	// for w := wCap; w >= 0; w-- {
	// 	if dp[n-1][w] >= 0 {
	// 		return dp[n-1][w]
	// 	}
	// }
	// return 0

	// 在最后一行中找到最大价值并返回
	result := math.MinInt
	for w := 0; w <= wCap; w++ {
		if dp[n-1][w] > result {
			result = dp[n-1][w]
		}
	}
	return result
}

func knapsackMaxVal2(weights []int, values []int, wCap int) int {
	if len(weights) != len(values) {
		panic("weights and values must have the same length")
	}
	if wCap < 0 {
		panic("wCap < 0")
	}
	if len(weights) == 0 {
		return 0
	}

	n := len(weights)
	// 创建状态转移表
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, wCap+1)
		for j := range len(dp[i]) {
			dp[i][j] = math.MinInt
		}
	}

	// 初始化第一行
	dp[0][0] = 0 // 不放，w:0, val: 0
	if weights[0] <= wCap {
		// 若 weights=[0,1] values= [-1,10], wCap=1, 则 expected:10, actual:9
		// dp[0][weights[0]] = values[0] // 放，w:weights[0], val: values[0]
		dp[0][weights[0]] = max(dp[0][weights[0]], values[0])
	}

	// 正向递推
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 注意，这里 w 是放入物品 i 前的重量，跟逆向时不同
			if dp[i-1][w] != math.MinInt { // 上一行对应单元格是否可达
				// 不放。比较当前单元格值跟来源单元格值，取大的
				dp[i][w] = max(dp[i][w], dp[i-1][w])
				if w+weights[i] <= wCap {
					// 放
					dp[i][w+weights[i]] = max(dp[i][w+weights[i]], dp[i-1][w]+values[i])
				}
			}
		}
	}

	// 逆向递推
	// 逆向状态转移方程：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i]]+values[i])
	// 注意，这里 w 语义跟正向递推时不同。
	// dp[i][w] 表示「考虑到物品 i 时，重量 w 是否可达，可达时最大价值」。到达它有两条路径：
	// dp[i-1][w] ----------不放--------> dp[i][w]
	// dp[i-1][w-weights[i]] -----放----> dp[i][w]
	// 只要任意一条路径可行，它就可达，就需要计算最大价值
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 有可能两条路径都是可达的
			// 上面的状态转移方程就是用下面两个 if 语句块实现的。逻辑上可以认为，如果两条路径都可达，
			// 状态转移方程：dp[i][w] = max(dp[i][w], dp[i-1][w], dp[i-1][w-weights[i]]+values[i])
			if dp[i-1][w] != math.MinInt {
				// 注意，dp[i][w] 可能已被其他树分支设置过了
				dp[i][w] = max(dp[i][w], dp[i-1][w])
			}
			// 注意确保 w 能够包含物品 i
			if w-weights[i] >= 0 && dp[i-1][w-weights[i]] != math.MinInt {
				dp[i][w] = max(dp[i][w], (dp[i-1][w-weights[i]] + values[i]))
			}
		}
	}

	result := math.MinInt
	for w := 0; w <= wCap; w++ {
		if dp[n-1][w] > result {
			result = dp[n-1][w]
		}
	}
	return result
}
