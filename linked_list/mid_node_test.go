package linkedlist

// https://leetcode.cn/problems/middle-of-the-linked-list/

// 算法思路：
// 1. 计算链表长度，找到中间位置
// 2. 遍历，找到中间节点
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	len := lenOfLinked(head)
	mid := len / 2
	curr := head
	index := 0
	for curr != nil {
		if index == mid {
			return curr
		}
		curr = curr.Next
		index++
	}
	return head
}

func lenOfLinked(head *ListNode) int {
	len := 0
	for p := head; p != nil; p = p.Next {
		len++
	}
	return len
}

// 算法思路：快慢指针
// 注意，节点数为奇数时，快指针结束于尾节点；为偶数时，快指针结束于尾节点的next，为 nil
func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	// for fast != nil && fast.Next != nil { // OK
	// Bug: 循环内的结束条件表达的是何时退出循环，而循环语句中的条件表达的是何时继续循环，
	// 因此不能直接将循环体中的结束条件拷贝作为循环继续条件
	//for fast == nil || fast.Next == nil {
	for {

		// 【编码技巧】若链表长度为偶数，fast 指向 tail.Next 时 slow 指向中间节点；
		// 若为奇数，fast 指向 tail 时 slow 指向中间节点。这意味着对于这两种情形，
		// for 循环的退出条件不同，因此就别试图在 for 结束声明那里一下子想清楚结束条件，
		// 先写个无限循环，再写循环体的基础核心逻辑或叫核心逻辑的骨架，再检查哪些地方为nil
		// 会引发错误，据此在循环体内或 for 结束声明那里补充结束条件

		// fast == nil: 节点为偶数；fast.Next == nil：节点为奇数
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		// 不需要保证 slow 非空，因为 fast 走在前面，fast 是 slow 的守卫
		slow = slow.Next
	}
	return slow
}

// 尝试把前一个版本中 for 循环体内的结束条件放在 for 声明中
func middleNode3(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 不需要保证 slow 非空，因为 fast 走在前面，fast 是 slow 的守卫
		slow = slow.Next
	}
	return slow
}
