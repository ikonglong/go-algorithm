package binary_tree

// https://leetcode.cn/problems/convert-bst-to-greater-tree
// 把二叉树转换为累加树

func convertBST(root *TreeNode) *TreeNode {
	// if root == nil {
	// 	return nil
	// }
	sum := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左-根-右
		inorder(root.Right)
		// 处理根节点
		sum += root.Val
		root.Val = sum
		inorder(root.Left)
	}
	inorder(root)
	return root
}
