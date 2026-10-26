package binary_tree

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	preorderTraverseFrom(root, &result)
	return result
}

func preorderTraverseFrom(root *TreeNode, result *[]int) {
	// root 就是当前节点
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preorderTraverseFrom(root.Left, result)
	preorderTraverseFrom(root.Right, result)
}
