package dp

// dp 逆向递推
// 状态转移表：dp [n][wCap+1]bool，n 是物品数量，单元格是列表示的状态是否可达到
// 状态转移方程：dp[i][w] = dp[i-1][w] || dp[i-1][w-weights[i]]
//
// weights 表示物品列表，同时也表示每个物品的重量
func knapsackMaxWt_(weights []int, wCap int) int {
	// 参数校验

	n := len(weights)
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, wCap+1)
	}

	// 初始化第一行
	dp[0][0] = true
	// 该死的！又忘记检查是否超重了
	if weights[0] > wCap {
		dp[0][weights[0]] = true
	}

	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			if dp[i-1][w] || (w-weights[i] >= 0 && dp[i-1][w-weights[i]]) {
				dp[i][w] = true
			}
		}
	}

	maxW := 0
	for w := wCap; w >= 0; w-- {
		if dp[n-1][w] && w > maxW {
			maxW = w
			break
		}
	}
	return maxW
}

// dp 正向递推
func knapsackMaxWt2_(weights []int, wCap int) int {
	// 参数校验

	n := len(weights)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, wCap+1)
	}

	// 初始化第一行
	dp[0][0] = true
	if weights[0] > wCap {
		dp[0][weights[0]] = true
	}

	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			if dp[i-1][w] {
				dp[i][w] = true
				if w+weights[i] <= wCap {
					dp[i][w+weights[i]] = true
				}
			}
		}
	}

	for w := 0; w <= wCap; w++ {
		if dp[n-1][w] {
			return w
		}
	}
	return 0
}
