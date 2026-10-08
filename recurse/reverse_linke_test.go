package recurse

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
// 反转链表

// 1. 反转链表 head 跟反转 head.next 解决思路一样
// 2. r(head): r(head.next)，再将反转得到的链表 tail.next 指向 head
func trainningPlan(head *ListNode) *ListNode {
	dummyHead, tail := reverse(head)
	tail.Next = nil
	return dummyHead.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		dummyHead := &ListNode{}
		tail := dummyHead
		return dummyHead, tail
	}
	dummyHead, tail := reverse(head.Next)
	tail.Next = head
	tail = head
	return dummyHead, tail
}
