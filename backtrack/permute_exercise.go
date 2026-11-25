package backtrack

import "slices"

// https://leetcode.cn/problems/permutations/
// https://leetcode.cn/problems/VvJkup/

func permute_(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	var backtrack func(k int, path []int)
	backtrack = func(k int, path []int) {
		if k == n {
			ans := make([]int, n)
			copy(ans, path)
			result = append(result, ans)
			return
		}

		for _, e := range nums {
			if slices.Contains(path, e) {
				continue
			}

			// 尝试一个选择
			path = append(path, e)
			backtrack(k+1, path)
			// 恢复状态，以尝试其他选择
			path = path[0 : len(path)-1]
		}
	}
	backtrack(0, nil)
	return result
}
