package binary_tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/

// 迭代实现
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 类似二分查找
	x := root
	for x != nil {
		if p.Val < x.Val && q.Val < x.Val {
			x = x.Left
		} else if p.Val > x.Val && q.Val > x.Val {
			x = x.Right
		} else {
			// 若 x 是 p/q，或者 p、q 分别在 x 的左右子树中，那么 x 必定是 LCA
			return x
		}
	}
	return nil
}

// 递归实现
func lowestCommonAncestor22(root, p, q *TreeNode) *TreeNode {
	// 若 root 是 p/q，或者 p、q 分别在 root 的左右子树中，那么 root 必定是 LCA
	if root == p || root == q ||
		(p.Val < root.Val && root.Val < q.Val) ||
		(q.Val < root.Val && root.Val < p.Val) {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		// 埋有低级 bug，导致 leetcode 报告超出内存限制
		// return lowestCommonAncestor22(root, p, q)
		return lowestCommonAncestor22(root.Left, p, q)
	} else {
		// return lowestCommonAncestor22(root, p, q)
		return lowestCommonAncestor22(root.Right, p, q)
	}
}

func lowestCommonAncestor23(root, p, q *TreeNode) *TreeNode {
	var binarySearch func(*TreeNode, *TreeNode, *[]*TreeNode)
	binarySearch = func(root *TreeNode, target *TreeNode, path *[]*TreeNode) {
		if root == nil {
			return
		}

		// 即使当前节点就是 p 或 q，也要添加至 path。否则，会遗漏这种情形：
		// 根节点为 p 或 q，另一个目标在左或右子树中
		*path = append(*path, root)
		if root == target {
			return
		} else if target.Val < root.Val {
			binarySearch(root.Left, target, path)
		} else {
			binarySearch(root.Right, target, path)
		}
	}

	pPath := make([]*TreeNode, 0)
	binarySearch(root, p, &pPath)
	qPath := make([]*TreeNode, 0)
	binarySearch(root, q, &qPath)
	var lca *TreeNode
	for i, j := 0, 0; i < len(pPath) && j < len(qPath); i, j = i+1, j+1 {
		if pPath[i] == qPath[j] {
			lca = pPath[i]
		} else {
			break
		}
	}
	return lca
}

func lowestCommonAncestor24(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var binarySearch func(*TreeNode, *TreeNode, *TreeNode)
	binarySearch = func(root, p, q *TreeNode) {
		if lca != nil { // 剪枝
			return
		}
		if root == nil {
			return
		}

		// 第一次尝试编写这个实现思路时卡住了，卡在不知道如何实现每次二分搜索两个对象。
		// 要一次搜索两个对象，那这两个对象就必须在根节点的同一侧子树中。只要先把所有不
		// 在同一侧子树中的情形处理掉，就可以继续递归二分搜索两个对象了。
		// 虽然这个推理说起来挺容易的，但我是在看了小争哥的迭代和递归实现后才梳理出来的

		// 若 root 是 p/q，或者 p、q 分别在 root 的左右子树中，那么 root 必定是 LCA
		if root == p || root == q ||
			(p.Val < root.Val && root.Val < q.Val) ||
			(q.Val < root.Val && root.Val < p.Val) {
			lca = root
			return
		}
		if p.Val < root.Val && q.Val < root.Val {
			binarySearch(root.Left, p, q)
		} else {
			binarySearch(root.Right, p, q)
		}
	}

	binarySearch(root, p, q)
	return lca
}
