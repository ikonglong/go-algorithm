package sort_

// https://leetcode.cn/problems/insertion-sort-list/
// 既然要返回排序后的链表，那就创建一个新链表，保存排序好的节点
// 关键点：
// - 遍历找插入位置时，找要挂在哪个节点后。
//   边界情形：要插在最前面，如何处理？
// 		a. 使用 dummy head，就把这种情形也转换为挂在某个节点后
//      b. 当没找到可插入位置时，把它插在头节点前，作为新的头节点

// 使用 dummy head，统一所有情形的处理方式
func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	// 外层循环遍历原链表
	for p := head; p != nil; {
		tmp := p.Next
		// 内层循环在有序新链表上找插入位置
		q := dummyHead
		// p > q.Next 时，虽然知道 p 肯定插入在 q.next 之后，但具体位置仍无法确定。
		// 只有遇到第一个 p <= q.Next，才确定 p 插入到 q 之后，q.next 之前
		for q.Next != nil && p.Val > q.Next.Val {
			q = q.Next
		}
		p.Next = q.Next
		q.Next = p

		// 移动 p
		p = tmp
	}
	return dummyHead.Next
}

// 不使用 dummy head，需要特殊处理要插在有序新链表的头节点之前的情形
func insertionSortList2(head *ListNode) *ListNode {
	var rHead *ListNode // result list head
	// 外层循环遍历原链表
	for p := head; p != nil; {
		tmp := p.Next
		// 内层循环在有序新链表上找插入位置
		if rHead == nil || p.Val < rHead.Val {
			p.Next = rHead
			rHead = p
		} else {
			q := rHead
			// p > q.Next 时，虽然知道 p 肯定插入在 q.next 之后，但具体位置仍无法确定。
			// 只有遇到第一个 p <= q.Next，才确定 p 插入到 q 之后，q.next 之前
			for q.Next != nil && p.Val > q.Next.Val {
				q = q.Next
			}
			p.Next = q.Next
			q.Next = p
		}

		// 移动 p
		p = tmp
	}
	return rHead
}
