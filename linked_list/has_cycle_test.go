package linkedlist

// https://leetcode.cn/problems/linked-list-cycle/

// 遍历链表。若节点指针在 map 中不存在，就放入；若存在，有环
func hasCycle_(head *ListNode) bool {
	aMap := make(map[*ListNode]struct{})
	p := head
	for p != nil {
		if _, found := aMap[p]; found {
			return true
		}
		aMap[p] = struct{}{}
		p = p.Next
	}
	return false
}

// 创建两个指针：slow 和 fast，slow 一次走一步，fast 一次走两步。
// 若有环，slow 和 fast 最终都会入环，在环中移动时也必定会相遇
func hasCycle2_(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast { // 相遇，必定有环
			return true
		}
	}
	return false
}
