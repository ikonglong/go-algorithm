package backtrack

// https://leetcode.cn/problems/subsets/

func subsets_(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	var backtrack func(int, []int)
	backtrack = func(k int, path []int) {
		// 有可行解了要退出
		if k == n {
			ans := make([]int, len(path))
			copy(ans, path)
			result = append(result, ans)
			return
		}
		// 递归终止条件

		// 不放 nums[k] 进组合
		backtrack(k+1, path)

		// 放 nums[k] 进组合
		path = append(path, nums[k])
		backtrack(k+1, path)
		// 恢复状态
		path = path[0 : len(path)-1]
	}
	backtrack(0, nil)
	return result
}
