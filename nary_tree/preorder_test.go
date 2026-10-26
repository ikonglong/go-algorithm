package nary_tree

// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
// N 叉树前序遍历

func preorder(root *Node) []int {
	result := make([]int, 0)
	var preorderFrom func(*Node)
	preorderFrom = func(root *Node) {
		// 最小子问题
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for _, child := range root.Children {
			preorderFrom(child)
		}
	}
	preorderFrom(root)
	return result
}
