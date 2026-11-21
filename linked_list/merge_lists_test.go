package linkedlist

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
// 合并两个训练计划

// 基于循环实现
func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}
	remaining := p1
	if p2 != nil {
		remaining = p2
	}
	tail.Next = remaining
	return dummyHead.Next
}

// 基于递归的实现
// 跟 recurse/merge_lists_test.go 重复
func trainningPlan2(l1 *ListNode, l2 *ListNode) *ListNode {
	var merge func(h1, h2 *ListNode) *ListNode
	merge = func(h1, h2 *ListNode) *ListNode {
		if h1 == nil {
			return h2
		} else if h2 == nil {
			return h1
		}

		var h *ListNode
		if h1.Val < h2.Val {
			h = h1
			h.Next = merge(h1.Next, h2)
		} else {
			h = h2
			h.Next = merge(h1, h2.Next)
		}

		// 脑残了犯这种低级错误
		// subHead := merge(h1.Next, h2.Next)
		// var h *ListNode
		// if h1.Val < h2.Val {
		// 	h = h1
		// 	h.Next = h2
		// } else {
		// 	h = h2
		// 	h.Next = h1
		// }
		// h.Next.Next = subHead

		return h
	}
	return merge(l1, l2)
}
