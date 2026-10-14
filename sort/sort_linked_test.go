package sort_

// https://leetcode.cn/problems/sort-list/solutions/492301/pai-xu-lian-biao-by-leetcode-solution/

// 采用非递归归并（自底向上的归并）实现
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n := lenOfLinked(head)
	step := 1
	for step < n {
		// 保存归并结果的新链表
		newDummyHead := &ListNode{}
		newTail := newDummyHead
		p := head // p 用来遍历 head 表示的链表
		for p != nil {
			// 每次循环，合并两个链表区间 [p, q], [q+1, r]

			// 先构造区间 [p, q]
			count := 1 // p 已被包含，因此计数为 1
			// 初始化 q 为 p，以便沿着链移动。
			// 为什么不初始化为 p.next？这样做区间就包含 q 了，跟 count=1 矛盾
			q := p
			for q != nil && count < step {
				q = q.Next
				count++
			}
			// 走到这里，q 可能为 nil。若为 nil，那么区间和计数都包含 null 值，但这点
			// 不会埋 bug，因为这表明待归并链表已经用尽了，将区间 [p, q] 挂在 tail 后，
			// 这一层归并就完成了，不会再使用 q 了

			// 若待归并链表已经用尽了，则将区间 [p, q] 挂在 tail 后，这一层归并就完成了
			if q == nil || q.Next == nil {
				newTail.Next = p
				break
			}

			// 若 q 后还有元素，则继续构造区间 [q+1, r]
			count = 1 // q+1 已被包含，因此计数为 1
			r := q.Next
			for r != nil && count < step {
				r = r.Next
				count++
			}
			// 注意，走到这里 r 可能为空。若为 nil，那么区间和计数都包含 null 值

			// 保存下一个 step 的起点，用于更新循环变量 p
			// 埋了 bug：迭代是从处理 p 开始的，但走到这里已经处理到 r 了
			// next := p.Next
			var next *ListNode
			if r != nil {
				next = r.Next
			}
			mergedHead, mergedTail := mergeSortedRanges(p, q, r)
			newTail.Next = mergedHead
			newTail = mergedTail
			p = next
		}

		// head: 待归并链表
		// newHead: 保存归并结果的新链表
		// 一层归并完成后，滚动 head, newHead
		head = newDummyHead.Next
		// 埋了 bug
		// head = newDummyHead
		step <<= 1 // step=2*step
	}

	return head
}

func lenOfLinked(head *ListNode) int {
	len := 0
	p := head
	for p != nil {
		len++
		p = p.Next
	}
	return len
}

func mergeSortedRanges(p, q, r *ListNode) (*ListNode, *ListNode) {
	dummyHead := &ListNode{}
	tail := dummyHead
	p1, p2 := p, q.Next
	sentinel1 := q.Next
	// 对于参数 r，调用方可能传入 nil。若 r 是 nil，则处理完其前驱节点后停止；
	// 若非 nil，则遇到 r.next 停止，不管其是否为 nil
	sentinel2 := (*ListNode)(nil)
	if r != nil {
		sentinel2 = r.Next
	}
	for p1 != sentinel1 || p2 != sentinel2 {
		var node *ListNode
		if p1 == sentinel1 { // l1 用尽
			node = p2
			p2 = p2.Next
		} else if p2 == sentinel2 { // l2 用尽
			node = p1
			p1 = p1.Next
		} else if p1.Val < p2.Val {
			node = p1
			p1 = p1.Next
		} else {
			node = p2
			p2 = p2.Next
		}

		tail.Next = node
		tail = tail.Next
	}
	// 合并结束后，若尾节点还保留着原链表中的 Next，就会形成环，导致无限循环
	tail.Next = nil
	return dummyHead.Next, tail
}

// 采用递归归并（自顶向下的归并）实现
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// 找到分割点
	mid := findMidNode(head)
	sorted1 := sortList2(head) // 何时结束
	sorted2 := sortList2(mid)
	return mergeTwoSorted(sorted1, sorted2)
}

func findMidNode(head *ListNode) *ListNode {
	// 参数链表必需至少有两个节点
	var prev, slow, fast *ListNode = nil, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	// 若只有一个节点，走到这里 prev 仍为 nil。
	// 可用 dummyhead 避免 prev 为 nil
	prev.Next = nil
	return slow
}

func mergeTwoSorted(h1, h2 *ListNode) *ListNode {
	var head, tail *ListNode
	p1, p2 := h1, h2
	for p1 != nil || p2 != nil {
		var node *ListNode
		if p1 == nil { // l1 用尽
			node = p2
			p2 = p2.Next
		} else if p2 == nil { // l2 用尽
			node = p1
			p1 = p1.Next
		} else if p1.Val < p2.Val {
			node = p1
			p1 = p1.Next
		} else {
			node = p2
			p2 = p2.Next
		}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}
	return head
}
