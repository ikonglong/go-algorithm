package dp

// weights 表示物品列表，同时也表示每个物品的重量
func knapsackMaxWt(weights []int, wCap int) int {
	n := len(weights) // 物品数量

	// 创建状态转移表格 dp：
	// 行数为 n，有多少个物品，就有多少行
	// 列数为 wCap+1，因为要用背包重量枚举(0..wCap)作为表格第二维的数组索引
	// 单元格表示列所示的背包重量是否到达了
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, wCap+1)
	}

	// 初始化状态转移表格 dp 的第一行
	// 物品 0，两种选择：放、不放
	dp[0][0] = true         // 不放重量为0，表示重量为0的列到达了
	if weights[0] <= wCap { // 这个检查很容易忘
		dp[0][weights[0]] = true
	}

	// 从上至下，逐行推导
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 检测上一行 w 列是否到达了。
			// 模拟的是决策树的当前层，从上一层的节点派生。
			// 未到达或不可达，意味着决策树中就不存在对应节点
			if dp[i-1][w] {
				// 从上一层的节点派生时，有两个选择：放、不放
				// 不放
				dp[i][w] = true
				if w+weights[i] <= wCap { // 这个检查很容易忘
					// 放。w 是上一行完成后背包的重量，即已放入物品的总重量
					// weights[r] 是当前物品的重量
					dp[i][w+weights[i]] = true
				}
			}
		}
	}

	for w := wCap; w >= 0; w-- {
		if dp[n-1][w] {
			return w
		}
	}

	return 0
}

// 再实现一遍。
// 将自上而下正向推导的实现，跟类似递归递推公式的从后向前逆向递归实现放在一起，对比理解。
// 注意，这里的逆向递推不是从最后一行开始向前逐行递推计算，而是已知第 i-1 行，计算
// 第 i 行时使用逆向递推公式计算，而不是正向推导计算
func knapsackMaxWt2(weights []int, wCap int) int {
	if len(weights) == 0 || wCap <= 0 {
		return 0
	}

	n := len(weights)
	// 创建状态转移表，每个单元格初始化为 false
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, wCap+1)
	}

	// 初始化第一行
	dp[0][0] = true         // 不放，w: 0
	if weights[0] <= wCap { // 放，w: weights[0]
		dp[0][weights[0]] = true
	}

	// 正向递推
	// for i := 1; i < n; i++ {
	// 	for w := 0; w <= wCap; w++ {
	// 		// 注意，这里 w 是放入物品 i 前的重量，跟逆向时不同
	// 		if dp[i-1][w] { // 上一行上方单元格是否可达
	// 			dp[i][w] = true // 不放物品i
	// 			if w+weights[i] <= wCap {
	// 				dp[i][w+weights[i]] = true // 放
	// 			}
	// 		}
	// 	}
	// }

	// 逆向递推
	// 逆向状态转移方程：dp[i][w] = dp[i-1][w] || dp[i-1][w-weights[i]]
	// 注意，这里 w 语义跟正向递推时不同。
	// dp[i][w] 表示「考虑到物品 i 时，重量 w 是否可达」。到达它有两条路径：
	// dp[i-1][w] ----------不放--------> dp[i][w]
	// dp[i-1][w-weights[i]] -----放----> dp[i][w]
	// 只要任意一条路径可行，结果就是可达。
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 注意确保 w 能够包含物品 i
			if dp[i-1][w] || (w-weights[i] >= 0 && dp[i-1][w-weights[i]]) {
				// 状态转移方程实现在哪里呢？就是 if 条件
				dp[i][w] = true
			}
		}
	}

	for w := wCap; w >= 0; w-- {
		if dp[n-1][w] {
			return w
		}
	}
	return 0
}
