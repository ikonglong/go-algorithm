package binary_tree

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
// 第一层按照从左到右的顺序记录
// 除第一层外每一层的记录顺序均与上一层相反。即第一层为从左到右，第二层为从右到左。

// 这个实现版本的算法思路或心智模型，请看下一个版本的注释
func decorateRecord3(root *TreeNode) [][]int { // leetcode name: decorateRecord
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	que := append([]*TreeNode(nil), root)
	levelCount := 0
	for len(que) != 0 {
		levelCount++
		// 外层循环每轮迭代开始时，队列中的元素就是将要处理的层的所有节点
		levelLen := len(que)
		levelArr := make([]int, levelLen)
		// levelArr := []int(nil)
		for i := range levelLen {
			node := que[i]
			if levelCount&1 == 1 {
				// 奇数层，从左到右填充
				levelArr[i] = node.Val
				// levelArr = append(levelArr, node.Val)
			} else {
				// 偶数层，从右到左填充
				levelArr[levelLen-1-i] = node.Val
				// levelArr = append([]int{node.Val}, levelArr...)
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		result = append(result, levelArr)
		que = que[levelLen:]
	}
	return result
}

// 埋有 bug：输入：[8,17,21,18,null,null,6]，输出：[[8],[21,17],[6,18]]，期望：[[8],[21,17],[18,6]]
//
// 原因分析：这里选择的思路是，处理一层节点时，根据层号的奇偶性决定是按索引从左到右，还是从右到左遍历。乍一看，
// 这个思路可行，但有一个很容易忽略的影响：处理一层节点的过程中，会将下一层节点放入遍历队列。这样，上一层节点的
// 遍历顺序就决定了下一层节点添加到队列的顺序。例如，处理第一层 [8]，奇数层，从左到右遍历，因此将 [8] 加入结果
// 列表，期间将 [17,21] 加入遍历队列；处理第二层 [17,21]，偶数层，从右到左遍历，因此将 [21,17] 加入结果
// 列表，期间将 [6,18] 加入遍历队列；处理第三层 [6,18]，奇数层，从左到右遍历，因此将 [6,18] 加入结果列表。
//
// 但第三层的顺序跟原始顺序不一样，上一层是从右到左遍历，致使按“从右到左顺序”将第三层节点添加到了遍历队列。如果
// 17、21 各有两个子节点，还会发现第三层节点的添加顺序也不是从左到右顺序，而是从右往左看上一层节点得到的
// [上层第一个节点的左、右子节点, 上层第二个节点的左、右子节点, 上层第 x 个节点的左右子节点]
//
// 要修复这个 bug，关键是按期望顺序生成结果列表时，不要影响下一层节点的添加顺序。如何做到？
// 这里的思路是通过改变遍历一层节点的顺序来生成期望的结果列表。我们可以不改变遍历顺序，通过改变填充结果列表的顺序
// 来生成期望结果：若是奇数层，从左向右填充；偶数层：从右向左填充
//
// 修复后的版本就是上一个版本

func decorateRecord32(root *TreeNode) [][]int { // leetcode name: decorateRecord
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	que := append([]*TreeNode(nil), root)
	levelCount := 0
	for len(que) != 0 {
		levelCount++
		// 外层循环每轮迭代开始时，队列中的元素就是将要处理的层的所有节点
		levelLen := len(que)
		levelArr := make([]int, 0)
		if levelCount&1 == 1 {
			// 奇数层，从左到右遍历
			for i := range levelLen {
				node := que[i]
				levelArr = append(levelArr, node.Val)
				if node.Left != nil {
					que = append(que, node.Left)
				}
				if node.Right != nil {
					que = append(que, node.Right)
				}
			}
		} else {
			// 偶数层，从右到左遍历
			for i := levelLen - 1; i >= 0; i-- {
				node := que[i]
				levelArr = append(levelArr, node.Val)
				if node.Left != nil {
					que = append(que, node.Left)
				}
				if node.Right != nil {
					que = append(que, node.Right)
				}
			}
		}

		result = append(result, levelArr)
		que = que[levelLen:]
	}
	return result
}

// 这个版本实现思路跟上一版本完全相同，只是提取出了便利函数，去除重复代码。
// 因此，它跟上个版本问题相同：
// 输入：[8,17,21,18,null,null,6]，输出：[[8],[21,17],[6,18]]，期望：[[8],[21,17],[18,6]]
func decorateRecord33(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	que := append([]*TreeNode(nil), root)
	levelCount := 0
	for len(que) != 0 {
		levelCount++
		// 外层循环每轮迭代开始时，队列中的元素就是将要处理的层的所有节点
		levelLen := len(que)
		levelArr := make([]int, 0)
		elemHandler := func(node *TreeNode) {
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		if levelCount&1 == 1 {
			// 奇数层，从左到右遍历
			traverseFromLToR(que, 0, levelLen, elemHandler)
		} else {
			// 偶数层，从右到左遍历
			traverseFromRToL(que, levelLen-1, 0, elemHandler)
		}

		result = append(result, levelArr)
		que = que[levelLen:]
	}
	return result
}

func traverseFromLToR(que []*TreeNode, start, end int, f func(node *TreeNode)) {
	for i := start; i < end; i++ {
		f(que[i])
	}
}

func traverseFromRToL(que []*TreeNode, start, end int, f func(node *TreeNode)) {
	for i := start; i >= end; i-- {
		f(que[i])
	}
}
