package binary_tree

import (
	"github.com/ikonglong/go-algorithm/common"
)

func postorderTraversal(root *common.TreeNode) []int {
	result := make([]int, 0)
	postorderTraversalRecurse(root, &result)
	return result
}

func postorderTraversalRecurse(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}

	postorderTraversalRecurse(root.Left, result)
	postorderTraversalRecurse(root.Right, result)
	*result = append(*result, root.Val)
}
