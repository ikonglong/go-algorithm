package linkedlist

// https://leetcode.cn/problems/lru-cache/

type LRUCache struct {
	size     int
	capacity int
	// 采用双向循环链表。双向使得删除中间节点简单
	// 为什么需要循环？因为 dummyHead、dummyTail 位于链表两端，它们就是两端的守卫，
	// 添加、删除节点操作都在它们之间进行，这样便可保护各种操作不那么容易出错。初始化时，
	// 要将它们两个链接起来，只能用循环链表。
	// 不用它们会面临什么问题？请看 linked_list/lru_bad_test.go
	dummyHead *DLinkedNode
	dummyTail *DLinkedNode
	// 为了确保 get 和 put 的平均时间复杂度为 O(1)
	// key 为缓存键
	nodeMap map[int]*DLinkedNode
}

// Constructor
func LRUCache(capacity int) LRUCache {
	dummyHead := &DLinkedNode{}
	dummyTail := &DLinkedNode{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead
	return LRUCache{
		capacity:  capacity,
		dummyHead: dummyHead,
		dummyTail: dummyTail,
		nodeMap:   make(map[int]*DLinkedNode),
	}
}

// 1. 在 nodeMap 中查找，未找到返回 -1
// 2. 找到节点，先将其删除，再将其添加为头节点
func (this *LRUCache) Get(key int) int {
	if n, found := this.nodeMap[key]; found {
		// 先删除，再添加为头
		// Bug: 若 addToHead 没有包含从 map 删除的操作，
		// 但 remove 包含了添加至 map 的操作，就会导致 bug
		this.remove(n)
		this.addAsHead(n)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// key 存在，更新值，删除，添加为头
	if n, found := this.nodeMap[key]; found {
		n.Val = value
		// Bug: 若 addToHead 没有包含从 map 删除的操作，
		// 但 remove 包含了添加至 map 的操作，就会导致 bug
		this.remove(n)
		this.addAsHead(n)
	} else {
		// key 不存在，先检查是否已满。
		// 若已满，则先淘汰最久访问的
		// 再创建新节点，添加为头，添加到 map，size+1
		if this.isFull() {
			this.remove(this.dummyTail.Prev)
		}
		n := &DLinkedNode{
			// Bug: 该死的！居然漏设置 key 了
			Key: key,
			Val: value,
		}
		this.addAsHead(n)
		// 不在这里做，统一交给 addAsHead 做
		// this.nodeMap[key] = n
	}
}

// 1. 先在 nodeMap 中查找
// 2. 未找到，直接返回
// 3. 找到，从链表和哈希表移除它
func (this *LRUCache) Del(key int) {
	// 先在map中查找
	// 存在，从双链中删除，从map删除，size-1
	// 不存在，什么也不做
	if n, found := this.nodeMap[key]; found {
		this.remove(n)
	}
}

// 对比下 linked_list/lru_bad_test.go 这两个方法的实现

// 若 addToHead 没有包含从 map 删除的操作，但 remove 包含了添加至 map 的操作，
// 就会导致 get、put 操作有 bug

func (this *LRUCache) addAsHead(node *DLinkedNode) {
	// 注意每个节点都有两个指针要设置，不要漏
	node.Next = this.dummyHead.Next
	node.Prev = this.dummyHead
	this.dummyHead.Next = node
	node.Next.Prev = node

	this.nodeMap[node.Key] = node
	this.size++
}

func (this *LRUCache) remove(node *DLinkedNode) {
	this.size--
	delete(this.nodeMap, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) isFull() bool {
	return this.size == this.capacity
}
