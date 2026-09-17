package linkedlist

// 题目分析：
// 从链表头到尾，依次是个位数、十位数、百位数等等
//
// 算法：
// 1. 创建一个结果列表，用来保存计算结果
// 2. 个位数相加。将两个链表头节点相加，若和大于等于 10，则向下一位进 1，否则进 0
// 3. 两个十位数、进位的数相加。若和大于等于十，则向下一位进 1，否则进 0
// 4. 以此类推，计算其他位，逐个添加到结果链表中。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummyHead := &ListNode{}
	current := dummyHead

	p1, p2 := l1, l2
	carry := 0 // 进位的数
	// 这行代码千万不能放在循环外。
	// 因为即使某个链表较短，也要将其对应的值置为 0
	// v1, v2 := 0, 0
	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next
	}

	// 这个 case 很容易遗漏。
	// 例如 5 加 5，两个链表同时遍历完，但产生了一个进位数 1
	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return dummyHead.Next
}

// 递归实现版本
//
// 算法：
//  1. 递归分解问题。将两个链表相加，分解为
//     a. 两个链表头节点相加，产出进位数
//     b. 移除头节点之后的子链表、进位数相加
//  2. 递归的基线条件：两个列表的头节点都为空
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	return addTwoNums(l1, l2, 0)
}

// 这个版本 OK，只是代码没那么简洁
func addTwoNums(head1 *ListNode, head2 *ListNode, carry int) *ListNode {
	if head1 == nil && head2 == nil {
		if carry == 0 { // 遗漏这种状况会导致高位为 0 时，仍然添加了一个节点
			return nil
		} else {
			// 将这个分支的逻辑合并到下面的主体逻辑中，就可以移除这个分支了。
			// 请看 addTwoNums2 实现
			return &ListNode{
				Val: carry,
			}
		}
	}

	// OK，但不够简洁
	// v1, v2 := 0, 0
	// var nextHead1, nextHead2 *ListNode
	// if head1 != nil {
	// 	v1 = head1.Val
	// 	nextHead1 = head1.Next
	// }
	// if head2 != nil {
	// 	v2 = head2.Val
	// 	nextHead2 = head2.Next
	// }
	// sum := (v1 + v2 + carry)

	// 简洁的版本。改进思路：
	// 既然目标是求 sum 或 val（结果节点的 val），就先定义这个变量，
	// 并为其设置合适的初始值（carry），过程中再不断累加
	sum := carry
	var nextHead1, nextHead2 *ListNode
	if head1 != nil {
		nextHead1 = head1.Next
		sum += head1.Val
	}
	if head2 != nil {
		nextHead2 = head2.Next
		sum += head2.Val
	}

	nextResult := addTwoNums(nextHead1, nextHead2, sum/10)
	return &ListNode{
		Val:  sum % 10,
		Next: nextResult,
	}
}

// OK，更简洁的版本
func addTwoNums2(head1, head2 *ListNode, carry int) *ListNode {
	if head1 == nil && head2 == nil && carry == 0 {
		return nil
	}

	result := &ListNode{}

	sum := carry
	var nextHead1, nextHead2 *ListNode
	if head1 != nil {
		sum += head1.Val
		nextHead1 = head1.Next
	}
	if head2 != nil {
		sum += head2.Val
		nextHead2 = head2.Next
	}

	result.Val = sum % 10

	// 如果控制流走到这里，可能是：head1 is nil, head2 is nil, carry != 0
	// 什么时候不走这个分支？head1 == nil && head2 == nil 为真就不走
	// 这其实是将 head1 is nil, head2 is nil, carry != 0 这个基线分支
	// 合并到了主逻辑中一起处理，代码确实更简洁了，但感觉不太好理解
	//
	// 下面的条件判断这样理解比较自然：head1、head2 中任何一个不为空，才能取它的 next

	if head1 != nil || head2 != nil {
		nextResult := addTwoNums2(nextHead1, nextHead2, sum/10)
		result.Next = nextResult
	}

	return result
}
