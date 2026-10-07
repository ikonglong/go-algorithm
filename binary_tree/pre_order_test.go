package binary_tree

import (
	"testing"

	"github.com/ikonglong/go-algorithm/common"
)

func preorderTraversal(root *common.TreeNode) []int {
	result := make([]int, 0)
	preorderTraversalRecurse(root, &result)
	return result
}

func preorderTraversalRecurse(root *common.TreeNode, result *[]int) {
	// root 就是当前节点
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preorderTraversalRecurse(root.Left, result)
	preorderTraversalRecurse(root.Right, result)
}

func TestPreorderTraversal(t *testing.T) {

}
