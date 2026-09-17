package linkedlist

import "fmt"

// import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "<nil>"
	}
	if n.Next == nil {
		return fmt.Sprintf("%d", n.Val)
	}
	return fmt.Sprintf("%d -> ", n.Val)
}

// func (n *ListNode) String() string {
// 	// return fmt.Sprintf("%v", *n) // dead loop
// 	return fmt.Sprintf("{%d, %v}", n.Val, *n.Next)
// }

func (h *ListNode) ValList() []int {
	if h == nil {
		return nil
	}

	vList := make([]int, 0, 16)
	n := h
	for n != nil {
		vList = append(vList, n.Val)
		n = n.Next
	}
	return vList
}

func NewLinked(list []int) *ListNode {
	dummyHead := &ListNode{}
	var prev, next *ListNode = dummyHead, nil
	for _, v := range list {
		next = &ListNode{Val: v}
		prev.Next = next
		prev = next
	}
	return dummyHead.Next
}

func lenOf(head *ListNode) int {
	count := 0
	p := head
	for p != nil {
		count++
		p = p.Next
	}
	return count
}
