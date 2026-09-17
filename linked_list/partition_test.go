package linkedlist

import (
	"slices"
	"testing"
)

// 算法思路
// 创建两个链表：before 和 after
// 1. before 用来添加小于 x 的节点
// 2. after 用来添加大于等于 x 的节点
// 注意：可以使用已有节点，也可以创建新节点。这需要跟面试官沟通

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var beforeDummyHead, afterDummyHead *ListNode = &ListNode{}, &ListNode{}
	beforeTail := beforeDummyHead
	afterTail := afterDummyHead

	p := head
	for p != nil {
		node := &ListNode{
			Val: p.Val,
		}
		if p.Val < x {
			beforeTail.Next = node
			beforeTail = beforeTail.Next
		} else {
			afterTail.Next = node
			afterTail = afterTail.Next
		}
		p = p.Next
	}

	if beforeDummyHead.Next == nil {
		return afterDummyHead.Next
	}

	beforeTail.Next = afterDummyHead.Next
	return beforeDummyHead.Next
}

// 不使用 dummy head 的版本
func partition_notUseDummyHead(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var beforeHead, beforeTail, afterHead, afterTail *ListNode

	p := head
	for p != nil {
		node := &ListNode{
			Val: p.Val,
		}
		if p.Val < x {
			if beforeHead == nil { // 表明添加第一个节点
				beforeHead = node
				beforeTail = beforeHead
			} else {
				beforeTail.Next = node
				beforeTail = beforeTail.Next
			}
		} else {
			if afterHead == nil {
				afterHead = node
				afterTail = afterHead
			} else {
				afterTail.Next = node
				afterTail = afterTail.Next
			}
		}
		p = p.Next
	}

	if beforeHead == nil {
		return afterHead
	}

	beforeTail.Next = afterHead
	return beforeHead
}

// 使用已有节点构建新链表，不创建已有节点
func partition_useExistingNodes(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var beforeHead, beforeTail, afterHead, afterTail *ListNode

	p := head
	for p != nil {
		next := p.Next
		p.Next = nil
		if p.Val < x {
			if beforeHead == nil { // 表明添加第一个节点
				beforeHead = p
				beforeTail = beforeHead
			} else {
				beforeTail.Next = p
				beforeTail = beforeTail.Next
			}
		} else {
			if afterHead == nil {
				afterHead = p
				afterTail = afterHead
			} else {
				afterTail.Next = p
				afterTail = afterTail.Next
			}
		}
		p = next
	}

	if beforeHead == nil {
		return afterHead
	}

	// 我犯过的最可饶恕，最低级的错误
	// beforeHead.Next = afterHead
	beforeTail.Next = afterHead
	return beforeHead
}

// 如果题目未明确是否要保持元素顺序，可以采用如下算法：
// 1. 这个算法就地将输入链表整理成期望的。基于此，初始时就需要将头节点看作「已整理好的链表」，它只包含头节点。
// 2. 创建两个指针，分别指向「已整理好的链表」头节点和尾节点。此时只有一个节点，就是原始链表的头节点
// 3. 遍历原始链表：
// 		如果小于 x，就将节点插在「已整理好的链表」头节点之前
//		如果大于等于 x，就添加在链表尾
//		每加入一个元素就要更新「已整理好的链表」的头节点或尾节点

func partition_inPlace(head *ListNode, x int) *ListNode {
	var pHead, pTail *ListNode = head, head

	// 这里让人费解的是，第一次循环时，node、pHead、pTail 都指向原始链表的头节点。
	// 如果其值小于 x，就将其链接到自己，即其下一个节点是其自身；如果大于等于 x 也如此。
	node := head
	for node != nil {
		next := node.Next
		node.Next = nil
		if node.Val < x {
			node.Next = pHead
			pHead = node
		} else {
			pTail.Next = node
			pTail = pTail.Next
		}
		node = next
	}

	// 如果 for 循环中没有 `node.Next = nil`，for 循环之后就需要下面这行
	// pTail.Next = nil

	return pHead
}

func TestPartition(t *testing.T) {
	tcs := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			head: NewLinked([]int{1, 4, 3, 2, 5, 2}),
			x:    3,
			want: NewLinked([]int{1, 2, 2, 4, 3, 5}),
		},
	}

	for _, tc := range tcs {
		// r := partition_useExistingNodes(tc.head, tc.x)
		r := partition_inPlace(tc.head, tc.x)
		if !slices.Equal(r.ValList(), tc.want.ValList()) {
			t.Errorf("partition_useExistingNodes(%v, %d) = %v, want: %v",
				tc.head.ValList(),
				tc.x,
				r.ValList(),
				tc.want.ValList())
		}
	}
}
