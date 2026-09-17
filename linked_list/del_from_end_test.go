package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := lenOf(head)
	nFromHead := len - (n - 1)

	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	current := head
	count := 0
	for current != nil {
		count++
		if count == nFromHead {
			prev.Next = current.Next
			break
		}
		prev = current
		current = current.Next
	}
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	removeNthFromEndRecurse(dummyHead, n)
	return dummyHead.Next
}

func removeNthFromEndRecurse(head *ListNode, n int) int {
	if head == nil {
		return 0
	}
	// 子链表的节点数
	k := removeNthFromEndRecurse(head.Next, n)
	if k > 0 && n == k {
		head.Next = head.Next.Next
	}
	// head 表示的链表的节点数
	return k + 1
}
