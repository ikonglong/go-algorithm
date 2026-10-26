package binary_tree

// https://leetcode.cn/problems/merge-two-binary-trees/description/
// 合并两个二叉树，可以分解为：合并左、右子树；再将合并后两个子树的根节点分别设置为当前节点的左、右子节点
// 最小子问题：合并两个空树
//
// 拿下述几个比最小子问题稍微复杂一点点的子问题走查验证：
//  1. 合并两个单节点树
//  2. 合并一个单节点树和一个空树
// 走查验证后，决定将子问题 2 也作为最小子问题之一单独处理

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	// 注意，条件 1 不满足意味着 root1 != nil or root2 != nil
	// 其中包含三种情形：
	// 1. root1 == nil && root2 != nil
	// 2. root1 != nil && root2 == nil
	// 3. root1 != nil && root2 != nil
	// 这里只需将前两种作为最小子问题之一单独处理
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil { // 条件 1 的 else：root1 != nil or root2 != nil
		root.Val = root2.Val
		return root2
	} else if root2 == nil {
		root.Val = root1.Val
		return root1
	}

	root.Val = root1.Val + root2.Val
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func mergeTrees_bug(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := &TreeNode{}
	if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
	} else if root1 != nil { // 条件 1 的 else：root1 == nil or root2 == nil
		root.Val = root1.Val
	} else {
		root.Val = root2.Val
	}
	// 埋有 bug：root1、root2 可能为 nil
	root.Left = mergeTrees_bug(root1.Left, root2.Left)
	root.Right = mergeTrees_bug(root1.Right, root2.Right)
	return root
}
