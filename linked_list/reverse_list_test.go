package linkedlist

// https://leetcode.cn/problems/reverse-linked-list/

// 创建一个结果链表，遍历输入链表，逐个头插到结果链表
func reverseList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := head
	for p != nil {
		next := p.Next
		p.Next = dummyHead.Next
		dummyHead.Next = p
		p = next
	}
	return dummyHead.Next
}

// 递归实现，定义全局结果链表
func reverseList2(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead
	var reverse func(*ListNode)
	reverse = func(head *ListNode) {
		if head == nil {
			return
		}
		reverse(head.Next)
		tail.Next = head
		tail = tail.Next
		// tail 就是原链表中的下个节点，它的 next 指向了 head，
		// 因此 head.next 就应该置为 nil，否则就成环了
		head.Next = nil
	}
	reverse(head)
	return dummyHead.Next
}

// 递归实现，不定义全局结果链表，而是返回子链表反转的结果
func reverseList3(head *ListNode) *ListNode {
	var reverse func(*ListNode) (*ListNode, *ListNode)
	reverse = func(head *ListNode) (*ListNode, *ListNode) {
		if head == nil {
			return nil, nil
			// dummyHead := &ListNode{}
			// tail := dummyHead
			// return dummyHead, tail
		}

		rHead, rTail := reverse(head.Next)
		// 让最小子问题返回 dummy head 和 tail 更优雅简洁
		if rTail == nil {
			rHead = &ListNode{
				Val: head.Val,
			}
			rTail = rHead
		} else {
			rTail.Next = head
			// 这句不能漏，否则 rTail 和 head 成环了
			head.Next = nil
			rTail = rTail.Next
		}
		return rHead, rTail
	}
	rHead, _ := reverse(head)
	return rHead
}
