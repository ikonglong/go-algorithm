package linkedlist

// https://leetcode.cn/problems/lru-cache/

// 因为没有用 dummyTail，所以维护 tail 指针、实现 remove(Node) 操作就很容易出错

type BadLRUCache struct {
	size     int
	capacity int
	// 采用双向链表。双向使得删除中间节点简单
	dummyHead *DLinkedNode
	// 注意，这不是 dummyTail
	tail *DLinkedNode
	// 为了确保 get 和 put 的平均时间复杂度为 O(1)
	// key 为缓存键
	nodeMap map[int]*DLinkedNode
}

// Constructor
func BadLRUCache(capacity int) BadLRUCache {
	dummyHead := &DLinkedNode{}
	return BadLRUCache{
		// Bug: 该死的！你居然漏设置 capacity 了
		dummyHead: dummyHead,
		tail:      dummyHead,
		nodeMap:   make(map[int]*DLinkedNode),
	}
}

// 1. 在 nodeMap 中查找，未找到返回 -1
// 2. 找到节点，先将其删除，再将其添加为头节点
func (this *BadLRUCache) Get(key int) int {
	if n, found := this.nodeMap[key]; found {
		// 先删除，再添加为头
		this.remove(n)
		this.addAsHead(n)
		return n.Val
	}
	return -1
}

func (this *BadLRUCache) Put(key int, value int) {
	// key 存在，更新值，删除，添加为头
	if n, found := this.nodeMap[key]; found {
		n.Val = value
		this.remove(n)
		this.addAsHead(n)
	} else {
		// key 不存在，创建新节点，添加为头，添加到 map，size+1
		n := &DLinkedNode{
			// Bug: 该死的！你居然漏设置 key 了
			Val: value,
		}
		this.addAsHead(n)
		this.nodeMap[key] = n
	}
}

// 1. 先在 nodeMap 中查找
// 2. 未找到，直接返回
// 3. 找到，从链表和哈希表移除它
func (this *BadLRUCache) Del(key int) {
	// 先在map中查找
	// 存在，从双链中删除，从map删除，size-1
	// 不存在，什么也不做
	if n, found := this.nodeMap[key]; found {
		this.remove(n)
	}
}

// 对比下 linked_list/lru_test.go 这两个方法的实现
// 如果没有下面这两个方法，维护 tail 指针就很容易出错

func (this *BadLRUCache) addAsHead(node *DLinkedNode) {
	// 如果链表为空
	if this.dummyHead.Next == nil {
		this.tail = node
	}

	// 注意每个节点都有两个指针要设置，不要漏
	// 这里漏设置了 node.Prev, 插入前 dummyHead.Next.Prev
	node.Next = this.dummyHead.Next
	this.dummyHead.Next = node
	this.size++
}

// 看看这里的问题，就知道为什么要有 dummyHead 和 dummyTail 了。
// 有了 dummyHead 和 dummyTail，如何将它们链接起来呢？
// 只能是 dummyHead.Next=dummyTail, dummyTail.Prev=dummyHead，
// 这样就构成了双向循环链表，它们两个就是链表两端的守卫，添加、删除节点
// 都在它们之间进行，这样就可保护各种操作不那么容易出错。
func (this *BadLRUCache) remove(node *DLinkedNode) {
	delete(this.nodeMap, node.Key)

	// 如果 node 是首节点，且没有 dummyHead，node.Prev 就会报错
	node.Prev.Next = node.Next
	// 如果 node 是尾节点，且没有 dummyTail，node.Next 就会报错
	node.Next.Prev = node.Prev

	// 如果删除的是尾节点，就要更新 tail
	// 然而，如果 node 是尾节点，就不可能执行到这里。
	// 因为没有 dummyTail，前一句中的 node.Next 就会报错
	if node.Next == nil {
		this.tail = node.Prev
	}

	this.size--
}
