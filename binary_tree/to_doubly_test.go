package binary_tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// 把二叉树转换为双向链表，顺序为中序，就地转换。left 指向前驱，right 指向后继

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dummyHead := &TreeNode{}
	tail := dummyHead
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// 处理 root
		tail.Right = root // right 用作 next 指针
		root.Left = tail  // left 用作 prev 指针
		tail = root
		inorder(root.Right)
	}

	inorder(root)
	// 构建双向循环
	dummyHead.Right.Left = tail
	tail.Right = dummyHead.Right
	return dummyHead.Right
}
