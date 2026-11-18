package binary_tree

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

func flatten(root *TreeNode) {
	dummyHead := &TreeNode{}
	tail := dummyHead
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		tail.Right = root // 将 right 指针用作 next
		tail = tail.Right // or tail = root
		//因为后面可能修改 root 到 right，root 到 left，先暂存
		left := root.Left
		right := root.Right
		root.Left = nil // 记得将 left 指针置空，否则结果就不是链表
		preorder(left)
		preorder(right)
	}
	preorder(root)
}
