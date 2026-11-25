package backtrack

// https://leetcode.cn/problems/subsets-ii/
//
// 需要关注是否有序

func subsetsWithDup_(nums []int) [][]int {
	opts := countFreq_(nums)

	result := make([][]int, 0)
	n := len(result)
	var backtrack func(int, []int)
	backtrack = func(k int, path []int) {
		if k == n {
			ans := make([]int, len(path))
			copy(ans, path)
			result = append(result, ans)
			return
		}

		// 当前阶段 k 要处理的数在原数组中出现了 c 次，
		// 当前阶段就有 c + 1 个选择
		for c := 0; c <= opts[k].count; c++ {
			for range c {
				path = append(path, opts[k].num)
			}
			backtrack(k+1, path)
			// 恢复状态
			path = path[0 : len(path)-c]
		}
	}

	backtrack(0, nil)
	return result
}
