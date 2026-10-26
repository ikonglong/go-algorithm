package binary_tree

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]

// 跟 binary_tree/node.go 中 ValList 实现相比较，这个实现优雅多了
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) > 0 { // 队列中没有节点了，表明已遍历完所有的树节点
		// 外层循环每轮迭代开始时，队列中的元素就是将要处理的层的所有节点
		levelLen := len(que)
		levelArr := make([]int, 0, levelLen)
		for i := range levelLen {
			node := que[i]
			// 将当前层节点值添加至结果列表
			levelArr = append(levelArr, node.Val)
			// 将下一层节点添加到队列，准备好下一个外层迭代要处理的
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		// if len(que) <= levelLen { // 不可能小于
		// 	que = nil
		// } else {
		// 	que = que[levelLen:]
		// }
		// 注意，对于 Go 语言，即使 que 长度为 k，
		// 创建切片时也可以这么写 que[k:]，不会报错
		que = que[levelLen:]
		result = append(result, levelArr)
	}
	return result
}

// 这个算法思路的心智模型，看下一个版本的注释
// 这个思路还有另一个变体，在每一层开始处插入层分隔元素。
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	levelArr := make([]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root, nil) // nil 为层分隔符
	// que example: [3,nil,9,20,nil,,,15,7,nil]
	for len(que) != 0 {
		head := que[0]
		if head == nil { // 层分隔符
			result = append(result, levelArr)
			levelArr = nil
			// 遍历到当前层的结束标识时，意味着下一层的所有节点已添加至遍历队列。因此，
			// 处理完当前层的结束标识后，添加一个结束标识到遍历队列，它就是下一层的结束标识
			//
			// 最后一层走到这里就处理完了，没有下一层了。若还为下一层添加结束标识，就会导致无限循环
			// 因此，要判断后面是否还有节点，确定是否是最后一层
			if len(que) > 1 {
				que = append(que, nil)
			}
		} else {
			levelArr = append(levelArr, head.Val)
			if head.Left != nil {
				que = append(que, head.Left)
			}
			if head.Right != nil {
				que = append(que, head.Right)
			}
			// 走到这里就为下一层添加结束标识是不对的
			// que = append(que, nil) // 添加层分隔符
		}

		// 将队首元素出队
		// if len(que) <= 1 {
		// 	que = nil
		// } else {
		// 	que = que[1:]
		// }

		// 注意，对于 Go 语言，即使 que 长度为 1，
		// 创建切片时也可以这么写 que[1:]，不会报错
		que = que[1:]
	}
	return result
}

// 这个思路没有 bug 的原因分析：
// 题目要求逐层遍历，但每一层的节点需要装入一个单独的数组。这就使得遍历过程中必须知道一层何时开始、何时结束。
// 这里的思路是，在将一层节点全部添加至遍历列表后，添加一个特殊元素，标识一层的结束。初始时，第一层只有
// 一个节点，结束标识元素添加得没问题。但在遍历队列的过程中，处理一层节点期间，将下一层节点添加至遍历队列时，
// 结束标识符设置得就不正确了。每将当前层一个节点的子节点添加至遍历列表后，就添加了一个下一层的结束标识元素，
// 但在那个位置下一层还没有结束呢。
//
// 如何修复？
// 假设当前层的层结束标识设置正确，那么根据当前层的层结束标识，就可以正确判断何时应该为下一层添加结束标识。
// 这个假设是成立的，因为第一层是硬编码设置的，可以认为一定正确。然后这里的遍历是一层循环，意味着每次循环只
// 处理一个元素，进而在遍历到当前层的结束标识时，已经将下一层所有节点添加至遍历队列了，因此处理完当前层的结束
// 标识后，向队列添加一个结束标识，这个标识就是下一层的正确结束标识了
//
// 上一个版本就是修复后的版本
func levelOrder_bug(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	levelArr := make([]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root, nil) // nil 为层分隔符
	// que example: [3,nil,9,20,nil,,,15,7,nil]
	for len(que) != 0 {
		head := que[0]
		if head == nil { // 层分隔符
			result = append(result, levelArr)
			levelArr = nil
		} else {
			levelArr = append(levelArr, head.Val)
			if head.Left != nil {
				que = append(que, head.Left)
			}
			if head.Right != nil {
				que = append(que, head.Right)
			}
			// 埋了 bug：对于 [3,9,20,null,null,15,7]，处理第二层（即添加[9,20]到结果列表）时，
			// 期望添加最后的 7 后只添加一个分隔符到队列，但实际是检测 9 的两个 null 子节点后添加了一个，
			// 添加 20 的 [15,7] 后又添加了一个。处理第三层（即添加[nil,15,7,nil]到结果列表）时，
			// 处理第一个 nil 时添加了一个空数组到结果列表；处理最后一个 nil 添加 [15,7] 到结果列表；
			// 处理 15、7 时，它们没有子节点，但仍分别添加了一个分隔符到队列，这导致下个迭代添加了两个
			// 空数组到结果列表。因此最后输出为：[[3],[9,20],[],[15,7],[],[]]
			//
			// 这么写表达的逻辑是：每次将一个节点的子节点添加到遍历队列，就添加一个层分隔符，而不是等下一层
			// 的所有节点都添加到队列后，再添加一个层分隔符
			que = append(que, nil) // 添加层分隔符
		}

		// 将队首元素出队
		// if len(que) <= 1 {
		// 	que = nil
		// } else {
		// 	que = que[1:]
		// }

		// 注意，对于 Go 语言，即使 que 长度为 1，
		// 创建切片时也可以这么写 que[1:]，不会报错
		que = que[1:]
	}
	return result
}
