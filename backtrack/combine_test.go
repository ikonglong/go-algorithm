package backtrack

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	// 按顺序处理 [0-n] 中的数字，每个数字一个决策阶段
	// 选择列表，每个阶段处理一个数字，两种选择：添加、不添加
	// 路径，已经添加到当前组合中的数字
	// 因为 k 已经被用作外层函数参数，因此用 t 表示阶段，
	// 同时表示当前阶段要处理的数字
	var backtrack func(t int, comb []int)
	backtrack = func(t int, comb []int) {
		if t == n+1 {
			if len(comb) == k {
				// 添加 comb 的快照
				result = append(result, append([]int(nil), comb...))
			}
			return
		}

		// 不添加
		backtrack(t+1, comb)

		// 添加
		comb = append(comb, t)
		backtrack(t+1, comb)
		//恢复
		comb = comb[:len(comb)-1]
	}

	backtrack(1, []int(nil))
	return result
}
