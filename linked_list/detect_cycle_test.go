package linkedlist

// 算法思路推导：
//
//  1. 检测是否有环。
//     创建两个指针 slow 和 fast。从头节点开始，slow 每次走一步，fast 每次走两步。
//     如果 fast 跟 slow 相遇了，则有环，否则无环，立即返回 null。
//
//  2. 假设从头节点到入环口长度为 k，即有 k 个节点。
//     环长度为大写 L，即包含 L 个节点。
//     对于环，入环口节点的索引为 0，节点索引依次递增。
//
//  3. 当 slow 走到入环口时，fast 走到距入环口 K 的位置。
//     距离入环口 K 的位置，意思是沿 next 方向（或前进方向）索引为 K 的节点处。
//     K = mod(2k-k, L) = mod(k, L)
//
//  4. 沿 next 方向观察，slow 落后 fast 的距离为 K，fast 落后 slow 的距离为 L - K。
//     fast 每移动一次追近一步。要追的距离是 L - K，因此需移动 L - K 次才能追完落后的距离。
//     注意，移动次数并不等于移动步数，要追的距离也并不等于追的过程中走的路程或距离。
//
//     那么【相遇】时：
//     fast 移动了 L - K 次，走了 2(L - K) 步，那么走的距离就是 2(L - K)
//     slow 也移动了 L - K 次，走了 (L - K) 步，那么走的距离就是 (L - K)
//     fast 和 slow 停在距离入环口 K 的位置。
//
//  5. 因为 K = mod(k, L)，所以 k = nL + K。
//     进而我们可以认为，当 slow 走到入环口时，fast 位于环上 k 位置。
//     即它在环上走了好 n 圈，最后一圈走完后，又向前走了 K 个节点。
//
//  6. 现在相遇点和链表头距离入环口都为 k。因此，让两个指针分别指向这两处，然后
//     再以相同的速度移动，当它们再次相遇时，相遇节点就是入环口
//
//
// 算法梳理：
// 1. 创建两个指针 slow 和 fast
// 2. slow 每次移动一步，fast 每次移动两步
// 3. 当它们相遇时，让 slow 指向链表头，fast 指向不变。
// 4. 两者同时以相同速度移动，每次一步。当再次相遇时，返回相遇节点

func detectCycle(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil { // 专门处理只有一个节点且没有环的情形
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// fast 比 slow 移动得快，因此 fast 是 slow 的守卫或哨兵，
	// 循环条件就不需要包含 slow 了
	//for slow != nil && fast != nil && fast.Next != nil {
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// 检测条件这样写，对于只有一个节点且没有环的情形就不正确。
	// 因为不满足循环条件，执行这个检测时又刚好满足了，就被误认为有环
	// if slow != nil && slow != fast { // 没有检测到环
	// 	return nil
	// }

	// 修订后的检测条件，专门针对循环退出条件设计的检测条件
	// 如果下面的条件满足，意味着循环只可能在两个指针相遇时退出
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

// 这个版本也 OK
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next

		if fast.Next == nil { // 表明无环
			return nil
		} else {
			fast = fast.Next.Next
		}

		if slow == fast {
			break
		}
	}

	if fast == nil { // 表明无环
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
