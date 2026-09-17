package linkedlist

import (
	"slices"
	"testing"
)

//https://leetcode.cn/problems/LGjMqU/?envType=problem-list-v2&envId=linked-list

// 先用快慢指针找到链表的中点。注意节点数为奇数时快指针到达末尾时，慢指针指向的是中点，
// 而偶数时，慢指针指向的是后一半的第一个节点。
// 再将后一半链表就地反转，最后合并两个链表。
func reorderList(head *ListNode) {
	prev := &ListNode{Next: head}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = prev.Next
		slow = slow.Next
		fast = fast.Next.Next
	}

	headOfSecondHalf := slow
	// 如果是奇数，fast 非 nil
	if fast != nil {
		headOfSecondHalf = slow.Next
		slow.Next = nil
	} else {
		prev.Next = nil
	}
	reversed := reverseInPlace(headOfSecondHalf)

	mergeTwoLists(head, reversed)
}

func reverseInPlace(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// [a1, a2, a3], [a5, a4]
// [a1, a2], [a4, a3]
func mergeTwoLists(head1 *ListNode, head2 *ListNode) {
	p1, p2 := head1, head2
	for p1 != nil && p2 != nil {
		next2 := p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = next2
	}
}

func TestReorderList(t *testing.T) {
	tcs := []struct {
		head *ListNode
		want *ListNode
	}{
		{NewLinked([]int{1, 2, 3, 4}), NewLinked([]int{1, 4, 2, 3})},
		{NewLinked([]int{1, 2, 3, 4, 5}), NewLinked([]int{1, 5, 2, 4, 3})},
	}

	for _, tc := range tcs {
		origin := tc.head.ValList()
		reorderList(tc.head)
		if !slices.Equal(tc.head.ValList(), tc.want.ValList()) {
			t.Errorf("reorderList(%v)=%v, want:%v", origin, tc.head.ValList(), tc.want.ValList())
		}
	}
}
