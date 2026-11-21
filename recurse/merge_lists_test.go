package recurse

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
// leetcode name: trainningPlan

// 采用递归实现
// 1. 原问题跟子问题的解决思路一样：
// 	  合并两个有序链表 跟 先把两个链表中较小的头节点摘除，合并剩余部分后再头插 思路类似
// 2. 最小子问题：任何一个链表为空，就返回另一个。两个链表可能同时为空

func mergeTwoSortedLinked(l1 *ListNode, l2 *ListNode) *ListNode {
	var mergeRecurse func(*ListNode, *ListNode) *ListNode
	mergeRecurse = func(head1, head2 *ListNode) *ListNode {
		if head1 == nil {
			return head2
		} else if head2 == nil {
			return head1
		}

		// subHead := mergeRecurse(head1.Next, head2.Next)
		// 比较两个链表的头节点。较小的节点，递归时传入其下一个节点；较大的节点，保持不动。
		// 例如：输入 [5] 和 [1,2,4]
		var mergedHead *ListNode
		if head1.Val < head2.Val {
			mergedHead = head1
			mergedHead.Next = mergeRecurse(head1.Next, head2)
		} else {
			mergedHead = head2
			mergedHead.Next = mergeRecurse(head1, head2.Next)
		}
		return mergedHead
	}
	return mergeRecurse(l1, l2)
}

func mergeTwoSortedLinked_bug(l1 *ListNode, l2 *ListNode) *ListNode {
	var mergeRecurse func(*ListNode, *ListNode) *ListNode
	mergeRecurse = func(head1, head2 *ListNode) *ListNode {
		if head1 == nil {
			return head2
		} else if head2 == nil {
			return head1
		}

		// 埋有 bug：递归时两个指针一起后移是不对的。
		// 例如：输入 [5] 和 [1,2,4]，结果为 [1,5,2,4]
		subHead := mergeRecurse(head1.Next, head2.Next)
		var mergedHead *ListNode
		if head1.Val < head2.Val {
			mergedHead = head1
			mergedHead.Next = head2
		} else {
			mergedHead = head2
			mergedHead.Next = head1
		}
		mergedHead.Next.Next = subHead
		return mergedHead
	}
	return mergeRecurse(l1, l2)
}
