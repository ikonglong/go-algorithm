package binary_tree

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// 从前序、后序序列构建二叉树

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return buildFromRange(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func buildFromRange(preorder []int, i, j int, postorder []int, p, r int) *TreeNode {
	// 该死，漏写了递归终止条件
	if i > j { // 此时 [i,j] 是空区间
		return nil
	}

	root := &TreeNode{
		Val: preorder[i],
	}
	// 必须处理这种情形，否则 i == j 时后面 preorder[i+1] 会访问到当前树范围之外的节点
	if i == j {
		return root
	}

	// 不管当前树有两个子树还是只有一个，先确定任意一个子树的大小，
	// 另一个子树的区间基于已知的子树大小计算。若另一个子树不存在，
	// 算出来的是空区间 [start, end] 且 start > end
	q := p
	// preorder[i+1] 就是其中一棵树的根节点值
	for postorder[q] != preorder[i+1] {
		q++
	}
	subTreeSize := q - p + 1

	// 注意，若当前树只有一个子树，根据前序和后序无法确定子树是左还是右。但题目认为设为左或右即可
	root.Left = buildFromRange(preorder, i+1, i+subTreeSize, postorder, p, q)
	// 该死，又犯低级错误了
	// root.Right = buildFromRange(postorder, i+subTreeSize+1, j, postorder, q+1, r)
	root.Right = buildFromRange(preorder, i+subTreeSize+1, j, postorder, q+1, r)
	return root
}

func buildFromRange_Bad(preorder []int, i, j int, postorder []int, p, r int) *TreeNode {
	// 该死，漏写了递归终止条件
	if i > j {
		return nil
	}

	// 尝试解决后面代码中 postorder[r-1]、preorder[i+1] 导致的两个问题：
	//   a) 数组访问越界
	//   b) 当前树只有一个节点时，它们会跑出当前树的范围。
	//        4
	//       / \
	//      1   3
	//      preorder:[4,1,3], postorder: [1,3,4]
	//      假设递归正在处理节点 3 这棵树，那么它在后序数组中的区间是 p = 1, r = 1，也就是只包含 postorder[1] == 3。
	//      这时 r-1 == 0，虽然 postorder[0] 是数组里的合法位置，但它对应的是节点 1，属于另一棵树，不属于当前节点 3 这棵树。
	//      所以问题不只是索引有没有超出整个数组，也要看它是否还在当前递归负责的区间里。
	//
	// r-1：当 p == r 时，后序区间只有一个节点，r-1 就在区间外。
	//   若 r == 0，还会直接变成负索引，导致越界；
	//   若 r > 0，虽然可能仍是合法的切片索引，但取到的是当前树之外的节点。
	// i+1：当 i == j 时，前序区间只有一个节点，i+1 就在区间外。
	//   若 i == len(preorder)-1，会直接越界；否则也可能读到当前树区间之外的节点。
	//

	// 尝试 1，根据实际越界情形进行防御
	// 埋有坑：只解决了越界。当前树只有一个节点时，仍会访问当前树之外的节点
	if r == 0 {
		return &TreeNode{
			Val: postorder[r],
		}
	}
	if i == j {
		return &TreeNode{
			Val: preorder[i],
		}
	}

	// 尝试 2，同时防御越界和访问当前树之外的节点
	// 仍然有坑：若当前树只有一个子树，递归处理左右子树时传入的区间就会错乱。详细原因分析见递归调用那里。
	if i == j {
		return &TreeNode{
			Val: preorder[i],
		}
	}
	if p == r {
		return &TreeNode{
			Val: postorder[r],
		}
	}

	// 找到根节点值
	rootVal := preorder[i]

	// 在后序中找到右子树的根。当前树的根是 postorder[r]，前一个就是右子树的根
	rightRootVal := postorder[r-1]
	// 拿右子树的根索引在前序中分割出左右子树节点区间
	// k 就是 preorder[i:j] 区间中右子树根所在位置
	k := i + 1
	for k <= j && preorder[k] != rightRootVal {
		k++
	}
	// 在前序中找到左子树的根。当前树的根是 preorder[i]，后一个就是左子树的根
	leftRootVal := preorder[i+1]
	// 拿左子树的根索引在后序中分割出左右子树节点区间
	// q 就是 postorder[p:r] 区间中左子树根所在位置
	q := p
	for q < r && postorder[q] != leftRootVal {
		q++
	}

	// 下面这段代码的 bug 分析：
	// i=0, j=1;  p=0, r=1
	// k=1     // 右子树根在前序中的位置
	// q=0    // 左子树根在后序中的位置
	// leftRoot = buildFromRange(preorder, i+1, k-1, postorder, p, q)
	// rightRoot = buildFromRange(preorder, k, j, postorder, q+1, r-1)
	//
	// 这段代码假设每个非叶节点都有左右两个孩子，但若当前树只有一个孩子，这段算法
	// 把同一个孩子同时当成左右子树的根，递归区间就错乱了：
	// 左：pre[1:0], post[0:0]  // [1:0] 空区间，两个区间大小不同
	// 右：pre[1:1], post[1:0]  // [1:0] 空区间，两个区间大小不同

	// 注意，若当前树只有一个子树，根据前序和后序无法确定子树是左还是右。但题目认为设为左或右即可
	leftRoot := buildFromRange(preorder, i+1, k-1, postorder, p, q)
	// 该死，应该传 q+1，而不是 q
	// rightRoot := buildFromRange(preorder, k, j, postorder, q, r-1)
	rightRoot := buildFromRange(preorder, k, j, postorder, q+1, r-1)
	return &TreeNode{
		Val:   rootVal,
		Left:  leftRoot,
		Right: rightRoot,
	}
}
