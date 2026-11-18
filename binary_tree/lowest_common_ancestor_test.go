package binary_tree

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 算法思路：
// 从 LCA 的定义转化入手，题目给的是一个直观易懂的描述，不是一个类似数学中那种形式化的，定量的定义。有没有这样的定义？
// 举例归纳，发现：
// 1. 若一个节点的左子树包含一个p/q，右子树也包含一个p/q，那么这个节点必定是p、q的LCA。【其他任何祖先都没有这个特点】
// 2. 若一个节点是p/q，其左/右子树也包含p/q，那么这个节点必定是p、q的LCA
// 原问题就转化为查找一个树，其左、右子树各包含几个 p/q
//
// 原问题就转化为【查找一棵树包含几个p/q】，尝试用递归解决这个问题。
// 如果知道左、右子树各包含几个p/q，再把根节点考虑上，就能算出一棵树包含几个 p/q。递归求解转化后问题的过程中，判断下是否是
// 左侧归纳的两种情况，就顺带把原问题解决了

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var dfs func(*TreeNode, *TreeNode, *TreeNode) int
	dfs = func(root, p, q *TreeNode) int { // or name as traverse
		if root == nil {
			return 0
		}
		lTargetCount := dfs(root.Left, p, q)
		if lca != nil {
			return 2
		}
		rTargetCount := dfs(root.Right, p, q)
		if lca != nil {
			return 2
		}
		rootTargetCount := 0
		if root == p || root == q {
			rootTargetCount = 1
		}
		// 左右子树各包含一个目标节点
		if rootTargetCount == 0 && lTargetCount == 1 && rTargetCount == 1 {
			// 别遗漏的这句
			lca = root
			return 2
		} else if rootTargetCount == 1 && (lTargetCount == 1 || rTargetCount == 1) {
			// 根节点是目标节点，且左或右子树包含一个目标节点
			lca = root
			return 2
		}
		return lTargetCount + rTargetCount + rootTargetCount
	}

	dfs(root, p, q)
	return lca
}
