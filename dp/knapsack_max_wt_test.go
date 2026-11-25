package dp

// dp 逆向递推
// 注意，逆向不是从最后一行开始向前逐行递推，而是思考方式的逆向，求 dp[i][w]，
// 先假设结果已知，再逆向思考是怎么计算出来的
//
// 状态转移表：dp [n][wCap+1]bool，n 是物品数量，单元格是列表示的状态是否可达到
// 状态转移方程：dp[i][w] = dp[i-1][w] || dp[i-1][w-weights[i]]
//
// 其中也包含自上而下的正向递推实现代码，用来跟逆向递推对比理解
//
// weights 表示物品列表，同时也表示每个物品的重量
func knapsackMaxWt(weights []int, wCap int) int {
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
	//      // 根据考虑了物品 i-1 后的重量 w 正向推导放、不放物品 i 后的重量。
	//      // 根据上一行的一个单元格，推导计算当前行两个单元格的值
	// 		if dp[i-1][w] { // 正上方单元格是否可达
	// 			// 不放
	// 			dp[i][w] = true
	//          // 放。先检查放了是否超重
	// 			if w+weights[i] <= wCap {
	// 				dp[i][w+weights[i]] = true
	// 			}
	// 		}
	// 	}
	// }

	// 逆向递推
	// 逆向状态转移方程：dp[i][w] = dp[i-1][w] || dp[i-1][w-weights[i]]
	// 注意，这里 w 语义跟正向递推时不同。
	// dp[i][w] 表示「考虑了物品 i 后，即决定放或不放后，重量 w 是否可达」。到达它有两条路径：
	// dp[i-1][w] ----------不放--------> dp[i][w]
	// dp[i-1][w-weights[i]] -----放----> dp[i][w]
	// 只要任意一条路径可行，结果就是可达。
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 根据上一行两个单元格的值，推导计算当前行【当前单元格】的值
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

// dp 正向递推
func knapsackMaxWt2(weights []int, wCap int) int {
	n := len(weights) // 物品数量

	// 创建状态转移表 dp：
	// 行数为 n，有多少个物品，就有多少行
	// 列数为 wCap+1，因为要用背包重量(0..wCap)作为表格第二维数组的索引
	// 单元格表示“列所示的”背包重量是否达到了
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, wCap+1)
	}

	// 初始化状态转移表 dp 的第一行
	// 物品 0，两种选择：放、不放
	dp[0][0] = true // 不放重量为 0，表示重量 0 可达了
	// 注意，weights[0] 是物品 0 的重量
	if weights[0] <= wCap { // 这个检查很容易忘
		dp[0][weights[0]] = true
	}

	// 从上至下，逐行推导
	for i := 1; i < n; i++ {
		for w := 0; w <= wCap; w++ {
			// 检测上一行 w 列是否到达了。
			// 模拟的是决策树的当前层，从上一层的节点派生。
			// 未达到或不可达，意味着决策树中就不存在对应节点
			if dp[i-1][w] {
				// 从上一层的节点派生时，有两个选择：放、不放
				// 不放
				dp[i][w] = true
				// 放。先检查放了是否超重，这个很容易忘
				if w+weights[i] <= wCap {
					// w 是上一行完成后背包的重量，即已放入物品的总重量
					// weights[i] 是当前物品的重量
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
