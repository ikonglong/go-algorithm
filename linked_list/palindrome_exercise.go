package linkedlist

// https://leetcode.cn/problems/palindrome-linked-list/

// 算法思路：先反转，再比较是否是回文。
// 优化：无需比较整个链表，只需比较到中间节点即可；那么反转时，是否能做到只反转到中间节点（实践发现不行）
// 先用快慢指针找到中间节点，注意节点数为偶数和奇数时，快指针最终位置不同
// 再反转链表
func isPalindrome_(head *ListNode) bool {
	// 至少两个节点才能用快慢指针
	if head == nil || head.Next == nil {
		return true
	}

	mid := findMidNode(head)
	rHead := reverseAndClone_(head)
	// bug
	// p1, p2 := rHead, head
	p1, p2 := head, rHead
	for p1 != mid.Next {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func findMidNode(head *ListNode) *ListNode {
	// 至少两个节点才能用快慢指针
	slow, fast := head, head
	for {
		// 偶数节点，fast 最终指向 tail.next
		// 奇数节点，fast 最终指向 tail
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// func reverseHalfAndClone(head, mid *ListNode) *ListNode {
func reverseAndClone_(head *ListNode) *ListNode {
	// 遍历+头插法反转必须走完整个链表，要只反转一半只能看看递归能否做到
	// 实践发现递归实现也无法做到只反转一半
	dummyHead := &ListNode{}
	tail := dummyHead
	var reverse func(*ListNode)
	reverse = func(head *ListNode) {
		if head == nil {
			return
		}

		reverse(head.Next)
		hCopy := &ListNode{
			Val: head.Val,
			// 别 copy next，否则 tail 和 hCopy 成环了
		}
		tail.Next = hCopy
		tail = tail.Next
	}
	reverse(head)
	return dummyHead.Next
}
