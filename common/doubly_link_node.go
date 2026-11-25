package common

// 双向链表节点
type DLinkedNode struct {
	Key  int // node 持有的业务对象的查找键
	Val  int
	Prev *DLinkedNode
	Next *DLinkedNode
}
