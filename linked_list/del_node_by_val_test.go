package linkedlist

import "math"

// https://leetcode.cn/problems/remove-linked-list-elements/

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  math.MaxInt,
		Next: head,
	}
	p := dummyHead
	// 【编码技巧】先写循环体逻辑，再检查哪些地方为nil会引发错误，据此补充循环结束条件
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			// 注意，这种情况不能移动 p，因为 p.Next 变了，需要再次检查 p.Next
		} else {
			p = p.Next
		}
	}
	return dummyHead.Next
}
