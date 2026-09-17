package linkedlist

// 题目分析：
// 1. 高位在链表头部，低位在链表尾部
// 2. 两个链表长度不一定相同
//
// 算法：
// 1. 遍历每个链表，获取其长度
// 2. 在较短链表的头部补零值节点，补充个数为两个链表的长度差值
// 3. 采用递归调用遍历两个链表。因为低位在链表尾部，所以递归到最低位时
//    才开始计算位上的数和进位数。在递归返回过程中，将高位结果节点插在低位的前面

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	len1 := lenOfII(l1)
	len2 := lenOfII(l2)

	head1, head2 := l1, l2
	if len1 < len2 {
		head1 = padList(l1, len2-len1)
	} else if len2 < len1 {
		head2 = padList(l2, len1-len2)
	}

	result, carry := addTwoNumsII(head1, head2)
	if carry > 0 {
		result = &ListNode{
			Val:  carry,
			Next: result,
		}
	}
	return result
}

func addTwoNumsII(head1, head2 *ListNode) (*ListNode, int) {
	if head1 == nil && head2 == nil {
		return nil, 0
	}

	result := &ListNode{}

	sum := 0
	var nextHead1, nextHead2 *ListNode
	if head1 != nil {
		sum += head1.Val
		nextHead1 = head1.Next
	}
	if head2 != nil {
		sum += head2.Val
		nextHead2 = head2.Next
	}

	// 虽然已经设置好了 nextHead1 和 nextHead2，但是设置这两个变量时
	// 先检查了 head1 和 head2 是否为 nil。因此，在使用这两个变量时
	// 也有必要添加语义相同的判断。但基线条件已经保证了如果两个链表都为空，
	// 不可能走到这里，因此这里没有必要加
	nextResult, carry := addTwoNumsII(nextHead1, nextHead2)

	sum += carry
	result.Val = sum % 10
	result.Next = nextResult

	return result, sum / 10
}

func lenOfII(head *ListNode) int {
	if head == nil {
		return 0
	}
	p := head
	count := 0
	for p != nil {
		count++
		p = p.Next
	}
	return count
}

func padList(head *ListNode, k int) *ListNode {
	p := head
	for k > 0 {
		p = &ListNode{
			Val:  0,
			Next: p,
		}
		k--
	}
	return p
}
