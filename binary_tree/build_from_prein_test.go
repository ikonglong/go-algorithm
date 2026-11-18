package binary_tree

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildFrom(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

// preorder[i] 是要构建的树的根节点
// preorder[i:j] 是树的前序顺序节点区间（闭区间），包含它自身
// preorder[p:r] 是树的中序顺序节点区间（闭区间），包含它自身
func buildFrom(preorder []int, i, j int, inorder []int, p, r int) *TreeNode {
	// 递归结束条件。[i,j]是闭区间，i==j时[i,j]表示叶子节点，仍需处理
	if i > j {
		return nil
	}

	// 先根据前序确定根节点
	rootVal := preorder[i]
	// 在中序序列中确定根节点位置，即找到分割点 q
	q := p
	for q <= r {
		if inorder[q] == rootVal {
			break
		}
		q++
	}

	// 分割点左边就是左子树区间 [i+1, i+leftSize] 和 [p+1, q-1]，
	// 右边就是右子树区间 [i+leftSize+1, j] 和 [q+1, r]
	leftSize := q - p
	// 埋了 bug：q 是根节点，已经排除。还要传入 p+1 意图排除根节点，那会儿我脑子是坏掉了么
	// leftRoot := buildFrom(preorder, i+1, i+leftSize, inorder, p+1, q-1)
	leftRoot := buildFrom(preorder, i+1, i+leftSize, inorder, p, q-1)
	rightRoot := buildFrom(preorder, i+leftSize+1, j, inorder, q+1, r)
	return &TreeNode{
		Val: rootVal,
		Left: leftRoot,
		Right: rightRoot,
	}
}
