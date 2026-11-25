package backtrack

// https://leetcode.cn/problems/combinations/

// 把一个具体的组合看作是一个背包
// 一个组合由几个元素组成，就有几个决策阶段
// 每个决策阶段决定：输入列表中的一个元素是否放入当前组合
// 注意，输入列表的大小可能比组合的大小大
func combine_(n int, k int) [][]int {
	result := make([][]int, 0)
	var backtrack func(int, []int)
	// step 当前阶段，也是当前要处理的元素
	backtrack = func(step int, path []int) {
		if len(path) == k {
			ans := make([]int, k)
			copy(ans, path)
			result = append(result, ans)
			return
		}

		// 递归终止条件
		if step == n+1 {
			return
		}

		// 不放入组合
		backtrack(step+1, path)

		// 放入组合
		path = append(path, step)
		backtrack(step+1, path)
		// 恢复状态
		path = path[0 : len(path)-1]
	}
	backtrack(1, nil)
	return result
}
