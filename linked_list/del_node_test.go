package linkedlist

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	node := head
	for node != nil {
		if node.Val == val {
			prev.Next = node.Next
			// node.Next = nil
			break
		}
		prev = node
		node = node.Next
	}
	return dummyHead.Next
}
