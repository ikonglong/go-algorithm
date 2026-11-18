package binary_tree

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 从中序、后序序列构建二叉树

func buildFromInPost(inorder []int, postorder []int) *TreeNode { // leetcode name: buildTree
	return buildFromInPostRange(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildFromInPostRange(in []int, i, j int, post []int, p, r int) *TreeNode {
	if i > j { // [i, j] 空区间
		return nil
	}

	root := &TreeNode{
		Val: post[r],
	}

	// 在中序序列上找到根节点，计算出左子树大小
	k := i
	for in[k] != root.Val {
		k++
	}
	leftTreeSize := k - i

	root.Left = buildFromInPostRange(in, i, k-1, post, p, p+leftTreeSize-1)
	root.Right = buildFromInPostRange(in, k+1, j, post, p+leftTreeSize, r-1)
	return root
}
