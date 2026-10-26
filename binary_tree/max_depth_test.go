package binary_tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 求一棵树的最大深度，可分解为：先求左右子树的最大深度，取较大值，再加一
// 最小值问题：root 为空返回 0

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
