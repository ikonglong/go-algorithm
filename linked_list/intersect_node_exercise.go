package linkedlist

// https://leetcode.cn/problems/intersection-of-two-linked-lists/

// 遍历链表。若当前节点的指针在 Map 中不存在，就放入 Map；若存在，则说明链表相交
func getIntersectionNode_exercise(headA, headB *ListNode) *ListNode {
	aMap := make(map[*ListNode]struct{})
	checkIntersection(headA, aMap)
	return checkIntersection(headB, aMap)
}

func checkIntersection(head *ListNode, aMap map[*ListNode]struct{}) *ListNode {
	p := head
	for p != nil {
		if _, found := aMap[p]; found {
			return p
		}

		aMap[p] = struct{}{}
		p = p.Next
	}
	return nil
}

// 1. 先计算两个链表的长度差值 d
// 2. 创建两个指针 p1、p2，分别指向两个链表的头节点。让较长链表的指针先移动 d 步
// 3. 然后同步遍历剩余部分，比较 p1、p2 是否相等，相等就相交了

func getIntersectionNode_exercise2(headA, headB *ListNode) *ListNode {
	aLen := LenOf(headA)
	bLen := LenOf(headB)

	// d := int(math.Abs(float64(aLen - bLen)))
	d := aLen - bLen
	p := headA
	if bLen > aLen {
		d = bLen - aLen
		p = headB
	}
	for c := 0; c < d; c++ {
		p = p.Next
	}

	var pA, pB *ListNode
	// bug: 上面 p 默认值为 pA。下面的语句导致长度相等时 pB 设为 pA
	//if aLen > bLen {
	if aLen >= bLen {
		pA = p
		pB = headB
	} else {
		pA = headA
		pB = p
	}
	for pA != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}

// 在上一版本的基础上尝试让代码简洁健壮些，不那么容易写出bug
func getIntersectionNode_exercise3(headA, headB *ListNode) *ListNode {
	aLen := LenOf(headA)
	bLen := LenOf(headB)

	// d := int(math.Abs(float64(aLen - bLen)))
	d := aLen - bLen
	longer := headA // 较长的链表
	if bLen > aLen {
		d = bLen - aLen
		longer = headB
	}
	p := longer
	for c := 0; c < d; c++ {
		p = p.Next
	}

	pA, pB := headA, headB
	if longer == headA {
		pA = p
	} else {
		pB = p
	}
	for pA != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}
