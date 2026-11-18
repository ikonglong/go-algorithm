package binary_tree

// https://leetcode.cn/problems/diameter-of-binary-tree/
// 首先，从左子树的最大直径、右子树的最大直径，无法推导出原树的最大直径。
// 接着，努力将原问题转化为可用递归求解的问题。直观理解，从左子树的最深节点走到
// 右子树的最深节点所经过的路径貌似就是最大直径。
// 这样就把求最大直径问题转化为求深度

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	// 树的深度
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lDepth := depth(root.Left)
		rDepth := depth(root.Right)
		// bug: 每次递归都会覆盖之前的最大直径
		// result = lDepth + rDepth
		result = max(lDepth+rDepth, result)
		return max(lDepth, rDepth) + 1
	}
	depth(root)
	return result
}
