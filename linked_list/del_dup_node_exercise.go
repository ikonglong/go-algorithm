package linkedlist

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	valMap := make(map[int]struct{})
	valMap[head.Val] = struct{}{}
	curr := head
	for curr != nil && curr.Next != nil {
		if _, found := valMap[curr.Next.Val]; found {
			curr.Next = curr.Next.Next
			// 因为 curr.next 变化了，所以不应该移动 curr
			// curr = curr.Next
		} else {
			valMap[curr.Next.Val] = struct{}{}
			curr = curr.Next
		}
	}

	return head
}
