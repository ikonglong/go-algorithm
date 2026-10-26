package binary_tree

// https://leetcode.cn/problems/invert-binary-tree/
// 翻转一个二叉树，可以分解为：先翻转其左右子树，再翻转其左右子节点
// 最小子问题：翻转一个空树
// 注意，翻转一个子树，子树的根节点除了左右孩子变化了外，其他都未变。
//
// 那下述几个比最小子问题复杂一点点的子问题，走查验证：
// 1. 叶子节点
// 2. 一个节点，只有左子节点且是叶子节点
// 3. 一个节点，只有右子节点且是叶子节点

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
