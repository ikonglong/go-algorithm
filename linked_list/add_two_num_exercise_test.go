package linkedlist

// https://leetcode.cn/problems/add-two-numbers/

// 容易遗漏处理的 case：最高位计算后，还要进 1
func addTwoNumbers_exercise(l1 *ListNode, l2 *ListNode) *ListNode {
	// sum := 0
	p1 := l1
	p2 := l2
	carry := 0    // 进位的数
	exponent := 0 // 10 的指数
	dummyHead := &ListNode{}
	tail := dummyHead
	for p1 != nil || p2 != nil {
		r := carry
		if p1 != nil {
			r += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			r += p2.Val
			p2 = p2.Next
		}

		val := 0
		if r >= 10 {
			// sum += (r - 10) * int(math.Pow10(exponent))
			val = r - 10
			carry = 1
		} else {
			// sum += r
			val = r
			// 这句不能漏
			carry = 0
		}
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
		exponent++

		// bug
		// p1 = p1.Next
		// p2 = p2.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
	// return sum
}
