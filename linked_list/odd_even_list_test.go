package linkedlist

// https://leetcode.cn/problems/odd-even-linked-list/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oddDummyHead := &ListNode{}
	oddTail := oddDummyHead
	evenDummyHead := &ListNode{}
	evenTail := evenDummyHead
	p := head
	count := 1
	for p != nil {
		if count&1 == 1 {
			oddTail.Next = p
			oddTail = oddTail.Next
		} else {
			evenTail.Next = p
			evenTail = evenTail.Next
		}

		p = p.Next
		count++
	}
	// 这句不能漏，否则可能成环
	evenTail.Next = nil
	oddTail.Next = evenDummyHead.Next
	return oddDummyHead.Next
}
