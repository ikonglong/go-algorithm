package linkedlist

// 1. 先把倒数第 k 个转换为正数第 n-k+1 个
// 2. 删除第 n-k+1 个节点
func removeNthFromEnd_(head *ListNode, n int) *ListNode {
	len := LenOf(head)
	// 倒数第 n 个的前一个节点
	kthFromStart := len - n
	dummyHead := &ListNode{Next: head}
	p := dummyHead
	count := 0
	for p != nil {
		if count == kthFromStart {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
		count++
	}
	return dummyHead.Next
}

// 借助递归思想，递进到尾节点后在归的过程中对节点计数。
// 回归到倒数第 n 个节点的前一个节点的调用中进行删除操作
func removeNthFromEnd2_(head *ListNode, n int) *ListNode {
	count := 0
	var traverse func(*ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}

		traverse(head.Next)
		// 可剪枝可不剪
		count++ // 计入当前节点
		// 回归到前一个节点的调用中了，做删除
		if count == n+1 {
			head.Next = head.Next.Next
		}
	}

	// 优雅简洁地处理要删除的是正数第一个节点
	dummyHead := &ListNode{Next: head}
	traverse(dummyHead)
	return dummyHead.Next
}
