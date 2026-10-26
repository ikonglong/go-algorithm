package nary_tree

// https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/
// 求一棵树的最大深度，可分解为：先求每个子树的最大深度，取较大值，再加一
// 最小值问题：root 为空返回 0

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	// 埋有 bug：maxDep 初始化成了 MinInt，对于叶子节点，因为没有子节点，所以循环一次也不执行。
	// 于是叶子节点会直接返回 `math.MinInt + 1`
	// maxDep := math.MinInt
	// 叶子节点没有子节点，因此其子树的最大深度应该为 0
	maxDep := 0
	for _, child := range root.Children {
		maxDep = max(maxDepth(child), maxDep)
	}
	return maxDep + 1
}
