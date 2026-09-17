package backtrack

import "slices"

// https://leetcode.cn/problems/permutations/
// 跟 backtrack/permute_test.go 题目相同，只是两个不同的 leetcode 题
//
// 需要关注输入集合是否有序，是否重复
func permute_2(nums []int) [][]int {
	result := make([][]int, 0)

	// 输入集合有多少个元素，就对应多少个决策阶段。
	// 选择列表：对于阶段 k，意味着已选择了k-1个元素，那么还剩下 n-k+1 种选择。
	//   不用一个单独的对象表示选择列表，遍历所有选择，过程中动态过滤出可用选择
	// 路径：已完成阶段添加到排列中的元素组成的排列
	var backtrack func(k int, path []int)
	backtrack = func(k int, path []int) {
		if k == len(nums) {
			result = append(result, append([]int(nil), path...))
			return
		}

		// 已选择了 k - 1 个元素，当前阶段还有 n - k + 1 种选择
		for _, v := range nums {
			if !slices.Contains(path, v) {
				path = append(path, v)
				backtrack(k+1, path)
				// 恢复
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []int(nil))
	return result
}
