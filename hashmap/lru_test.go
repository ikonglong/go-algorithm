package hashmap

type LRUCache struct {
	size      int
	capacity  int
	dummyHead *DLinkedNode // dummy head
	dummyTail *DLinkedNode // dummy tail
	nodeMap   map[int]*DLinkedNode
}

func Constructor(capacity int) LRUCache {
	h := &DLinkedNode{
		Key: -1,
		Val: -1,
	}
	t := &DLinkedNode{
		Key: -1,
		Val: -1,
	}
	h.Next = t
	// bug：之前遗漏了这句
	t.Prev = h
	return LRUCache{
		capacity:  capacity,
		dummyHead: h,
		dummyTail: t,
		nodeMap:   make(map[int]*DLinkedNode),
	}
}

// 1. 在 nodeMap 中查找。未找到返回 -1
// 2. 找到节点，先将其移除，再将其添加为头节点
func (this *LRUCache) Get(key int) int {
	// 添加如下检查有助于发现 size 相关的逻辑是否正确
	if this.size == 0 {
		return -1
	}
	if node, found := this.nodeMap[key]; found {
		this.remove(node)
		this.addAsHead(node)
		// bug：漏了这句，导致找到也没有返回节点值，始终返回 -1
		return node.Val
	}
	return -1
}

// 1. 先在 nodeMap 中查找
// 2. 未找到，再检查是否已满，若已满，先删除最久访问的节点。创建新节点，添加为头节点，添加到哈希表
// 3. 找到，更新节点数据，将其移除，再将其添加为头节点
func (this *LRUCache) Put(key int, value int) {
	if node, found := this.nodeMap[key]; found {
		node.Val = value
		this.remove(node)
		this.addAsHead(node)
	} else {
		if this.IsFull() {
			this.remove(this.dummyTail.Prev)
			// bug：遗漏了这两句。我认为应该把这两句放入 remove 中
			// delete(this.nodeMap, key)
			// this.size--
		}
		node := &DLinkedNode{
			Key: key,
			Val: value,
		}
		this.addAsHead(node)
	}
}

// 1. 先在 nodeMap 中查找
// 2. 未找到，直接返回
// 3. 找到，从链表和哈希表移除它
func (this *LRUCache) Del(key int) {
	// 添加如下检查有助于发现 size 相关的逻辑是否正确
	if this.size == 0 {
		return
	}
	if node, found := this.nodeMap[key]; found {
		this.remove(node)
		// bug：遗漏了这两句。方法注释里也漏了。我认为应该把这两句放入 remove 中
		// delete(this.nodeMap, key)
		// this.size--
	}
}

// 得益于头、尾节点前后各有一个守卫节点，这里可以将三种情形的删除逻辑
// 给统一、简化：
// 1. 删除中间节点
// 2. 删除尾节点
// 3. 删除仅剩的一个节点
func (this *LRUCache) remove(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	// node.Next.Prev = node.Prev
	// node.Prev.Next = node.Next

	delete(this.nodeMap, node.Key)
	this.size--
}

// or addAtHead
func (this *LRUCache) addAsHead(node *DLinkedNode) {
	node.Next = this.dummyHead.Next
	this.dummyHead.Next.Prev = node
	this.dummyHead.Next = node
	node.Prev = this.dummyHead

	this.nodeMap[node.Key] = node
	this.size++
}

func (this *LRUCache) IsFull() bool {
	return this.size == this.capacity
}

// 双向链表
type DLinkedNode struct {
	Key  int // 对应 node 持有的业务对象的查找键
	Val  int
	Prev *DLinkedNode
	Next *DLinkedNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
