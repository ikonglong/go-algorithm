package binary_tree

// https://leetcode.cn/problems/binode-lcci/

func convertBiNode(root *TreeNode) *TreeNode {
	dummyHead := &TreeNode{}
	tail := dummyHead
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// 把 right 指针用作 next
		tail.Right = root
		tail = root
		root.Left = nil
		inorder(root.Right)
	}
	inorder(root)
	return dummyHead.Right
}
