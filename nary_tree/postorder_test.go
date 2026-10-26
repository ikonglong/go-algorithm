package nary_tree

// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/
// N 叉树后续遍历

func postorder(root *Node) []int {
	result := make([]int, 0)
	var postorderFrom func(*Node)
	postorderFrom = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			postorderFrom(child)
		}
		result = append(result, root.Val)
	}
	postorderFrom(root)
	return result
}
