package nary_tree

// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	que := make([]*Node, 0)
	que = append(que, root)
	for len(que) != 0 {
		// 外层循环每轮迭代开始时，队列中的元素就是将要处理的层的所有节点
		levelLen := len(que)
		levelArr := make([]int, levelLen)
		for i := range levelLen {
			node := que[i]
			levelArr[i] = node.Val
			que = append(que, node.Children...)
		}
		result = append(result, levelArr)
		// 这句容易遗漏
		// 注意，对于 Go 语言，即使 que 长度为 k，
		// 创建切片时也可以这么写 que[k:]，不会报错
		que = que[levelLen:]
		que = que[levelLen:]
	}
	return result
}
