package linkedlist

import (
	"math"
)

// 算法：
// 1. 遍历两个链表，获得其长度和尾节点
// 2. 比较尾节点。如果尾节点不同（比较尾节点引用是否指向相同节点，而不是比较节点值），立即返回，两个链表无交点。
//    为什么比较尾节点就能判断是否相交呢？
//    因为题目所说的相交，意思是两个链表共用从相交节点到尾节点的所有节点。
// 3. 使用两个指针，分别指向两个链表的头部
// 4. 将较长链表的指针向前移动，移动步数为链表长度的差值
// 5. 现在同时遍历两个链表，直到两个指针指向相同节点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	var tA, tB *ListNode

	lenA, lenB := 0, 0
	for pA != nil {
		lenA++
		if pA.Next == nil {
			tA = pA
		}
		pA = pA.Next
	}
	for pB != nil {
		lenB++
		if pB.Next == nil {
			tB = pB
		}
		pB = pB.Next
	}

	if tA != tB {
		return nil
	}

	pA, pB = headA, headB
	if lenA > lenB {
		for i := 1; i <= lenA-lenB; i++ {
			pA = pA.Next
		}
	} else if lenB > lenA {
		for i := 1; i <= lenB-lenA; i++ {
			pB = pB.Next
		}
	}

	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tA, lenA := getTailAndSize(headA)
	tB, lenB := getTailAndSize(headB)

	if tA != tB {
		return nil
	}

	var shorter, longer *ListNode
	if lenA < lenB {
		shorter = headA
		longer = headB
	} else {
		shorter = headB
		longer = headA
	}

	kthNode := getKthNode(longer, int(math.Abs(float64(lenA-lenB))))
	p1, p2 := kthNode, shorter

	// OK，但不够简洁的版本
	// for p1 != nil && p2 != nil {
	// 	if p1 == p2 {
	// 		return p1
	// 	}
	// 	p1 = p1.Next
	// 	p2 = p2.Next
	// }
	// return nil

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

func getTailAndSize(head *ListNode) (*ListNode, int) {
	count := 1 // 下面的循环中，尾节点未被计数，因此初始值为 1
	p := head
	for p.Next != nil {
		count++
		p = p.Next
	}
	return p, count
}

// k 从零开始数
func getKthNode(head *ListNode, k int) *ListNode {
	// 有 bug，且啰嗦的版本
	// count := 0
	// p := head
	// for p != nil {
	//  p = p.Next
	// 	count++
	// 	if count == k {
	// 		return p
	// 	}
	// }
	//
	// 如果两个链表完全相同，k 等于 0，这行代码就是 bug
	// panic(fmt.Errorf("unexpected error: len > k(=%d)", k))

	// OK，但还不够简洁的版本
	// count := 0
	// p := head
	// for count < k && p != nil {
	// 	p = p.Next
	// 	count++
	// }
	// return p

	// 更简洁的版本
	p := head
	for k > 0 && p != nil {
		p = p.Next
		k--
	}
	return p
}
