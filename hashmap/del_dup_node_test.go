package hashmap

// https://leetcode.cn/problems/remove-duplicate-node-lcci/
// 题意分析：
// 有重复节点链表跟有环链表有些类似，感觉容易搞混。有环是指后面的节点链接至前面的节点。
// 重复节点，指的是节点的值有重复，不可能是两个不同的节点是同一个节点，这自相矛盾

// 算法思路：
// 1.创建一个 Map，用来保存节点指针；创建一个新链表，用来保存结果
// 2.遍历链表，若当前节点在 map 中不存在，就添加至 map 和结果链表；若存在，则跳过
//
// 其他思路，看：linked_list/del_duplicate_test.go

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	aMap := make(map[int]struct{})
	newDummyHead := &ListNode{}
	newTail := newDummyHead
	p := head
	for p != nil {
		if _, found := aMap[p.Val]; found {
			p = p.Next
			continue
		}

		next := p.Next
		// 如果不这么做，那么对于输入[1,2,3,3,2,1]，
		// 第一个节点就加入了结果链表，使得结果链表跟原始链表完全相同
		p.Next = nil
		aMap[p.Val] = struct{}{}
		newTail.Next = p
		newTail = p
		p = next
	}
	return newDummyHead.Next
}
