package binary_tree

// https://leetcode.cn/problems/list-of-depth-lcci/
// 把树的每一层节点从左到右构建成一个链表

func listOfDepth(tree *TreeNode) []*ListNode {
	result := make([]*ListNode, 0)
	que := make([]*TreeNode, 0)
	// 初始化时，先把第一层节点添加到队列
	que = append(que, tree)
	for len(que) != 0 {
		levelSize := len(que)
		levelDummyHead := &ListNode{}
		levelTail := levelDummyHead

		// for i := range levelSize {
		for _ = range levelSize {
			// 有点作弊，未使用 queue 接口，直接当数组使用的
			// elem := que[i]
			// 模拟队首元素出队
			elem := que[0]
			que = que[1:]

			levelTail.Next = &ListNode{
				Val: elem.Val,
			}
			levelTail = levelTail.Next

			// 别忘了展开
			if elem.Left != nil {
				que = append(que, elem.Left)
			}
			if elem.Right != nil {
				que = append(que, elem.Right)
			}
		}
		// que = que[levelSize:]
		result = append(result, levelDummyHead.Next)
	}
	return result
}
