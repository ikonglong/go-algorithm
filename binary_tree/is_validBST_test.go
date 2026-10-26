package binary_tree

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	isValid := true
	var minMax func(*TreeNode) (int, int)
	minMax = func(root *TreeNode) (min, max int) {
		// 初始化 min、max，可简化处理 left 或 right 为空的情形
		// 若 left 和 right 都为空，min、max 都应该等于根节点的值
		// 若 left 为空，最小值为 min；若 right 为空，最大值为 max
		min = root.Val
		max = root.Val
		if root.Left != nil {
			leftMin, leftMax := minMax(root.Left)
			if isValid == false { // 剪枝，提早返回
				// 此时返回什么都不会影响问题的正确求解了
				// return math.MinInt, math.MaxInt
				return min, max
			}
			if leftMax >= root.Val {
				isValid = false
				return min, max
			}
			// 这句容易漏
			min = leftMin
		}
		if root.Right != nil {
			rightMin, rightMax := minMax(root.Right)
			if isValid == false { // 剪枝
				return min, max
			}
			if root.Val >= rightMin {
				isValid = false
				return min, max
			}
			// 这句容易漏
			max = rightMax
		}
		return min, max
	}

	minMax(root)
	return isValid
}
