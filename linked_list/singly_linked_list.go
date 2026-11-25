package linkedlist

// https://leetcode.cn/problems/design-linked-list/

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	p := this.head
	i := 0
	for p != nil {
		if i == index {
			return p.Val
		}
		p = p.Next
		i++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{
		Val:  val,
		Next: this.head,
	}
	this.head = newHead
	// 这句容易漏
	if this.head.Next == nil {
		this.tail = newHead
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val: val,
	}

	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.Next = node
		this.tail = node
	}

	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		// 按题目要求，这里不应该报错
		// panic(fmt.Errorf("illegal index %d", index))
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	i := 0
	p := this.head
	for p != nil {
		if i == index-1 {
			node := &ListNode{Val: val}
			node.Next = p.Next
			p.Next = node
			break
		}
		p = p.Next
		i++
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		// 按题目要求，这里不应该报错
		// panic(fmt.Errorf("illegal index %d", index))
		return
	}

	// 这句很容易漏
	this.size--

	if index == 0 {
		this.head = this.head.Next
		// 这个判断更健壮，不会因为 size 的改变而不正确
		if this.head == nil {
			// Bug: 如果先修改了 size 就不能用下面这个判断
			// if index == this.size-1 {
			this.tail = nil
		}
		return
	}

	i := 0
	p := this.head
	for p != nil {
		if i == index-1 {
			p.Next = p.Next.Next
			break
		}

		p = p.Next
		i++
	}

	// 这个判断更健壮，不会因为 size 的改变而不正确
	if p.Next == nil {
		// Bug: 如果先修改了 size 就不能用下面这个判断
		// if index == this.size-1 {
		this.tail = p
	}
}
