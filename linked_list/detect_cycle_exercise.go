package linkedlist

// https://leetcode.cn/problems/linked-list-cycle-ii/

//  1. 检测是否有环：若无环，则立即返回 null；若有环，则继续找到入环口。
//     创建两个指针：slow 和 fast，slow 每次走一个节点，fast 每次走两个节点；
//     它们都从链表头出发，若相遇，则有环，否则无环
//
//  2. 在环上相遇时，相遇位置到入环口的路程，跟链表头到入环口的路程相等，因此，
//     让 slow 从链表头开始，fast 从相遇位置开始，每次移动一个节点。当它们再次相遇时，
//     相遇位置就是入环节点。所谓【路程相等】，意思是：此时若让两个指针分别从相遇位置和
//     链表头出发，每次移动一个节点，那么环上指针可以一直转圈圈，以等待跟另一个指针相遇。
func detectCycle_exercise(head *ListNode) *ListNode {
	// 1. 检测是否有环
	hasCycle := false
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}

	// 2. 找到入环口
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
