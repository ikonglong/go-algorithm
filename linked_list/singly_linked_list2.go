package linkedlist

// https://leetcode.cn/problems/design-linked-list/

// 借助虚拟头节点简化实现

type MyLinkedList2 struct {
	dummyHead *ListNode
	tail      *ListNode
	size      int
}

func Constructor2() MyLinkedList2 {
	node := &ListNode{}
	return MyLinkedList2{
		dummyHead: node,
		tail:      node,
	}
}

func (this *MyLinkedList2) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	i := 0
	p := this.dummyHead.Next
	for p != nil {
		if i == index {
			return p.Val
		}
		p = p.Next
		// 又写漏了
		i++
	}

	return -1
}

func (this *MyLinkedList2) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: this.dummyHead.Next,
	}
	this.dummyHead.Next = node
	// 这局不能漏，但我又写漏了
	if this.dummyHead.Next.Next == nil {
		this.tail = node
	}
	this.size++
}

func (this *MyLinkedList2) AddAtTail(val int) {
	this.tail.Next = &ListNode{
		Val: val,
	}
	this.tail = this.tail.Next
	this.size++
}

func (this *MyLinkedList2) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	// Bug: 如果插入在 0 位置，index-1=-1，而 i 永远不可能为负，没插入
	// 下面这两行就没有利用上虚拟投节点，跟没有它一样
	// i := 0
	// p := this.dummyHead.Next
	i := -1
	p := this.dummyHead
	for p != nil {
		if i == index-1 {
			node := &ListNode{
				Val:  val,
				Next: p.Next,
			}
			p.Next = node
			break
		}

		p = p.Next
		i++
	}

	this.size++
}

func (this *MyLinkedList2) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	// i 是 prev.Next 的下标，而非 prev 的下标
	// 	i := 0
	// 	prev := this.dummyHead
	// 	for prev != nil {
	// 		if i == index {
	// 			prev.Next = prev.Next.Next
	// 			break
	// 		}
	//
	// 		prev = prev.Next
	// 		i++
	// 	}

	// prev 的下标
	i := -1
	prev := this.dummyHead
	for prev != nil {
		if i == index-1 {
			prev.Next = prev.Next.Next
			break
		}

		prev = prev.Next
		i++
	}

	if prev.Next == nil {
		this.tail = prev
	}

	this.size--
}
