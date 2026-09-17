package backtrack

// https://leetcode.cn/problems/TVdhkn/description/
// 对比 backtrack/permute_test.go
//
// 这个实现是如何做到不生成重复组合的呢？
// 每层递归只处理 k 位置的元素，并且只有两个选择：选或不选。
// 随后，k 只会递增，不会回头。因此，元素始终按照它们在原始集合中的顺序加入到
// path 中，因此不会生成 [2, 1]，因为处理完 2 后不能返回去选择 1。
func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	// 决策阶段：处理原集合中的一个元素 nums[k]
	// 选择列表：是否将 nums[k] 添加到当前组合（combination）
	// 路径：保存已添加到组合中的元素
	var backtrack func(nums []int, k int, path []int)
	backtrack = func(nums []int, k int, path []int) {
		// 结束条件
		if k == len(nums) {
			// 注意：需要复制 path 的快照，不能直接将 path 添加到结果列表中，
			// 因为 path 在其他递归树分支中还会被修改
			result = append(result, append([]int(nil), path...))
			return
		}

		// 做选择
		// 这个问题中不存在要剪枝的情形

		// 不添加
		backtrack(nums, k+1, path)

		// 添加
		path = append(path, nums[k])
		backtrack(nums, k+1, path)
		// 恢复状态
		path = path[:len(path)-1]
	}

	backtrack(nums, 0, []int{})
	return result
}
