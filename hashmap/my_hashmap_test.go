package hashmap

// https://leetcode.cn/problems/design-hashmap/

const SLOT_COUNT int = 3535

type MyHashMap struct {
	slots []*hashListNode
}

// Constructor
func NewMyHashMap() MyHashMap {
	return MyHashMap{
		slots: make([]*hashListNode, SLOT_COUNT),
	}
}

// 若对应 slot 为空，则为该 slot 创建链表，将 KV 对添加至末尾
// 若 slot 不为空，则查找对应节点。
//  1. 若找到，更新节点值
//  2. 若未找到，添加至链表末尾
func (this *MyHashMap) Put(key int, value int) {
	slotIdx := this.hash(key)
	slot := this.slots[slotIdx]
	pair := Pair{
		key:   key,
		value: value,
	}
	if slot == nil {
		head := addAtTail(slot, pair)
		this.slots[slotIdx] = head
	} else if n := findByKey(slot, key); n != nil {
		n.data.value = value
	} else {
		addAtTail(slot, pair)
	}
}

func (this *MyHashMap) Get(key int) int {
	slotIdx := this.hash(key)
	slot := this.slots[slotIdx]
	n := findByKey(slot, key)
	if n != nil {
		return n.data.value
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	slotIdx := this.hash(key)
	slot := this.slots[slotIdx]
	head := remove(slot, key)
	this.slots[slotIdx] = head
}

func (this *MyHashMap) hash(key int) int {
	return key % SLOT_COUNT
}

type Pair struct {
	key   int
	value int
}

type hashListNode struct {
	data Pair
	next *hashListNode
}

func addAtTail(head *hashListNode, data Pair) *hashListNode {
	node := &hashListNode{
		data: data,
	}
	if head == nil {
		head = node
		return head
	}

	p := head
	// 为什么这里循环条件不确保 p != nil？
	// 因为循环之前已确保 p 初始时不为 nil，那么 p 从首元素走到尾元素都不可能为 nil
	for p.next != nil {
		p = p.next
	}
	p.next = node

	return head
}

func findByKey(head *hashListNode, key int) *hashListNode {
	p := head
	for p != nil {
		if p.data.key == key {
			return p
		}
		p = p.next
	}
	return nil
}

func remove(head *hashListNode, key int) *hashListNode {
	if head == nil {
		return nil
	}
	p := head
	// 为什么这里循环条件不确保 p != nil？
	// 因为循环之前已确保 p 初始时不为 nil，那么 p 从首元素走到尾元素都不可能为 nil
	for p.next != nil && p.next.data.key != key {
		p = p.next
	}
	// 未找到。注意未找到中包含头结点是目标节点的情形
	if p.next == nil {
		if head.data.key == key {
			head = head.next
			return head
		} else {
			return head
		}
	}
	// 找到。注意找到中包含尾节点是目标节点的情形，但不需要特殊处理，因为是单向链表
	p.next = p.next.next
	return head
}
