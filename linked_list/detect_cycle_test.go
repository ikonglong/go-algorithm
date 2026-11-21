package linkedlist

// https://leetcode.cn/problems/linked-list-cycle-ii/

// 算法思路推导：
//
//  1. 检测是否有环，若无，立即返回 null。若有，继续找到入环口
//     创建两个指针 slow 和 fast。都从头节点开始，slow 每次走一个节点，fast 每次走两个节点。
//     如果 fast 跟 slow 相遇了，则有环，否则无环。
//
//  2. 假设从链表头到入环口长度为 x，即有 x 个节点。
//     环长度为 L，即包含 L 个节点。对于环，入环口节点的索引为 0，节点索引依次递增。
//
//  3. 当 slow 走到入环口时，fast 走到距入环口 K 的位置。此时，slow、fast 走过的路程分别是 x、2x。
//     K = mod(2x-x, L) = mod(x, L)
//
//     距离入环口 K 的位置，意思是沿 next 方向（或前进方向）索引为 K 的节点处。
//     沿 next 方向观察，slow 落后 fast 的路程为 K，fast 落后 slow 的路程为 L - K。
//     fast 每移动一步追近一个节点。要追的路程是 L - K，因此需移动 L - K 步才能追完落后的路程。
//     注意，要追的路程或距离【并不等于】追的过程中走过的路程或距离。
//
//  4. 从 slow 走到入环口到在环中【相遇】时，
//     fast 移动了 L - K 次，走了 2(L - K) 个节点，那么走的路程就是 2(L - K)
//     slow 也移动了 L - K 次，走了 (L - K) 个节点，那么走的路程就是 (L - K)
//     fast 和 slow 停在距离入环口 L-(L-K) = K 的位置。
//
//  5. K = mod(x, L) => x = nL + K
//     相遇时到入环口的距离为 K，从链表头到入环口长度为 x = nL + K。
//     此时，如果让两个指针分别从链表头、环上相遇位置开始走，每次移动一步，当它们再次相遇时所在节点就是入环节点。
//     因为环上指针可以一直转圈圈，以等待跟另一个指针相遇，所以可以认为第一次相遇时它们到入环口的【路程】都是 x。
//
// 算法总结：
//
//  1. 检测是否有环：若无环，则立即返回 null；若有环，则继续找到入环口。
//     创建两个指针：slow 和 fast，slow 每次走一个节点，fast 每次走两个节点；
//     它们都从链表头出发，若相遇，则有环，否则无环
//
//  2. 在环上相遇时，相遇位置到入环口的路程，跟链表头到入环口的路程相等，因此，
//     让 slow 从链表头开始，fast 从相遇位置开始，每次移动一个节点。当它们再次相遇时，
//     相遇位置就是入环节点。所谓【路程相等】，意思是：此时若让两个指针分别从相遇位置和
//     链表头出发，每次移动一个节点，那么环上指针可以一直转圈圈，以等待跟另一个指针相遇。

func detectCycle(head *ListNode) *ListNode {
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

// 这个实现版本代码，啰嗦且不易读，不推荐
func detectCycle2(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil { // 专门处理只有一个节点且没有环的情形
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// fast 比 slow 移动得快，fast 是 slow 的守卫或哨兵，因此循环条件就不需要包含 slow 了
	//for slow != nil && fast != nil && fast.Next != nil {
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// 检测条件这样写，对于只有一个节点且没有环的情形就不正确，
	// 因为不满足循环条件，但又刚好满足了这个检测，就被误认为有环
	// if slow != nil && slow != fast { // 没有检测到环
	// 	return nil
	// }

	// 修订后的检测条件，专门针对循环退出条件设计的检测条件
	// 如果下面的条件不满足，意味着循环只可能在两个指针相遇时结束
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
}
