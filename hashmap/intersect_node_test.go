package hashmap

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 题意分析：
// 1. 列表节点的值可能重复
// 2. 相交的意思是，它们在内存中指向相同的位置，不是两个节点值相同
//
// 是否可用双指针解？

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aMap := make(map[*ListNode]struct{})
	p := headA
	for p != nil {
		aMap[p] = struct{}{}
		// 又把这句漏了
		p = p.Next
	}

	p = headB
	for p != nil {
		if _, found := aMap[p]; found {
			return p
		}
		// 又把这句漏了
		p = p.Next
	}
	return nil
}
