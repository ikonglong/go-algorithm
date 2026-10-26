package linkedlist

// 算法思路：
// 1. 创建两个指针：previous 和 current，分别表示前一个节点和当前节点。
//    	如果链表参数命名是 n 或 node，不是 head，可以将参数名用作当前节点。
//    	初始时，previous 为空，current 指向头节点。
//    再创建一个 HashMap 用来检测是否重复
// 2. 通过 current 遍历整个链表
// 		如果当前节点值在 map 中不存在，则将其放入 map，向前移动 previous 和 current
// 		否则，将 previous 链接至 current 的下一个节点（删除了 current），
// 		再移动 current 至下一个节点

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 将 previous 也初始化为 head 可能导致理解不顺畅
	// prev, current := head, head
	var prev, current *ListNode = nil, head
	set := make(map[int]struct{})
	for current != nil {
		if _, found := set[current.Val]; found {
			// 此时，previous 不需要移动，
			// 因为它已经连接至移动后的 current 了
			prev.Next = current.Next
			// current = current.Next
		} else {
			set[current.Val] = struct{}{}
			prev = current
			// current = current.Next
		}
		current = current.Next
	}

	return head
}

// 不借助额外缓冲区删除重复节点的算法。算法思路：
// 1. 创建两个指针：current 和 runner。
//		current：遍历整个链表，即访问链表的每个节点
// 		runner：检查 current 的后继节点中是否有跟 current 重复
//			如果检查后继节点的过程中，runner 直接指向被查看节点，又如何实现删除呢？
// 			runner 不能指向被查看节点，它总是指向被查看节点的前驱节点。也就是说，
// 			实际检查的不是 runner，而是 runner.next

func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current != nil {
		// 检查后继节点是否有重复
		runner := current
		for runner.Next != nil {
			if runner.Next.Val == current.Val {
				// 注意，此时删除的不是 runner，而是 runner.next
				// runner 就是 previous
				// 千万别写作：
				// runner = runner.Next.Next
				runner.Next = runner.Next.Next
				// 然后 runner 需要移动吗？
				// 不需要，因为 runner 链接至了新 next，它仍然扮演 previous。
				// 下一次循环继续检查 runner.next
			} else {
				// 不重复，runner 就需要移动
				runner = runner.Next
			}
		}

		current = current.Next
	}

	return head
}
