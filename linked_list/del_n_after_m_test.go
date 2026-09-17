package linkedlist

import (
	"fmt"
	"slices"
	"testing"
)

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	p := head
	for {
		fmt.Printf("start p: %d\n", p.Val)
		// p 指向保留片段最后一个节点
		for i := 1; i <= m-1 && p != nil; i++ {
			p = p.Next
		}
		if p == nil {
			break
		}
		fmt.Printf("moved to the last retained, p: %d\n", p.Val)

		// p2 指向删除片段的最后一个节点
		p2 := p
		fmt.Printf("p2: %v\n", p2)
		for i := 1; i <= n && p2 != nil; i++ {
			p2 = p2.Next
		}

		// 删除
		if p2 == nil { // 移动到末尾了
			p.Next = nil
			break
		} else {
			p.Next = p2.Next
			p2.Next = nil
			// 移动 p 到下一个保留片段的第一个节点
			p = p.Next
		}
		fmt.Printf("moved to the first of next retained part, p: %d\n", p.Val)
	}
	return head
}

// deleteNodes 跟 deleteNodes2, deleteNodes3 相比之下，前一种实现方式更容易犯错，
// 因为弄了两个指针，有时候 P 向前移动，有时候 P2 向前移动，所以在实现逻辑中要检测两个
// 指针是否到达了末尾。
//
// 后两种实现的思路更好，因为一个指针在前，一个指针在后，只需要关注前面那个指针是否到达
// 了末尾就够了，前面的指针是后面的指针的守卫，实现起来就没那么容易犯错。

func deleteNodes2(head *ListNode, m int, n int) *ListNode {
	var lastM, current *ListNode = nil, head
	for current != nil {
		for i := 1; i <= m-1 && current != nil; i++ {
			// 跟 deleteNodes3 相比，这里没有一起移动 lastM，实现就复杂些
			current = current.Next
		}
		if current == nil {
			break
		}
		lastM = current

		for i := 1; i <= n+1 && current != nil; i++ {
			current = current.Next
		}
		if current == nil {
			lastM.Next = nil
			break
		} else {
			lastM.Next = current
			lastM = current
		}
	}
	return head
}

func deleteNodes3(head *ListNode, m int, n int) *ListNode {
	var lastM, current *ListNode = nil, head
	for current != nil {
		mCount, nCount := m, n
		// 将 lastM 移动到保留部分的最后一个节点
		for current != nil && mCount > 0 {
			// 这个阶段 lastM 总是跟 current 一起移动，且总是指向 current 的前驱
			lastM = current
			current = current.Next
			mCount--
		}

		// 将 current 移动到下一个保留部分的第一个节点
		for current != nil && nCount > 0 {
			current = current.Next
			nCount--
		}

		// 删除 current 前面的 n 个节点
		lastM.Next = current
	}
	return head
}

func TestDeleteNodes(t *testing.T) {
	testCases := []struct {
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{NewLinked([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}), 2, 3, NewLinked([]int{1, 2, 6, 7, 11, 12})},
		{NewLinked([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}), 1, 3, NewLinked([]int{1, 5, 9})},
		{NewLinked([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}), 3, 1, NewLinked([]int{1, 2, 3, 5, 6, 7, 9, 10, 11})},
		{NewLinked([]int{9, 3, 7, 7, 9, 10, 8, 2}), 1, 2, NewLinked([]int{9, 7, 8})},
		{NewLinked([]int{6, 3, 5, 6, 2, 8, 9, 2, 3, 4}), 2, 1, NewLinked([]int{6, 3, 6, 2, 9, 2, 4})},
	}

	for _, tc := range testCases {
		got := deleteNodes3(tc.head, tc.m, tc.n)
		gotList := got.ValList()
		wantList := tc.want.ValList()
		if !slices.Equal(gotList, wantList) {
			t.Errorf("deleteNodes(%v, %d, %d)=%v, want: %v", tc.head.ValList(), tc.m, tc.n, gotList, wantList)
		}
	}
}
