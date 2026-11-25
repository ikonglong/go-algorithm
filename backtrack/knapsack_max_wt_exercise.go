package backtrack

// 每个物品对应一个处理阶段
// 每个阶段有两个选择：放、不放
// path: 题目没要求计算最终放入背包的物品列表，因此不需要路径
// 但需要记录累计重量
func knapsackMaxWt_(weights []int, wCap int) int {
	maxW := 0
	n := len(weights)

	var backtrack func(k int, cw int)
	backtrack = func(k int, cw int) {
		// cw == wCap 剪枝
		if cw == wCap || k == n {
			if cw > maxW {
				maxW = cw
			}
			return
		}

		// 不放
		backtrack(k+1, cw)

		// 放。先检查是否超重
		if cw+weights[k] > wCap {
			return
		}
		backtrack(k+1, cw+weights[k])
		// 没有改变全局和局部状态，因此不需要恢复状态
	}

	backtrack(0, 0)
	return maxW
}
