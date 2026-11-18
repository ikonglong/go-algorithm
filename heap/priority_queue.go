package heap

import (
	"container/heap"
	"fmt"
)

// 堆的底层数据结构
type HeapList[T any] struct {
	items []T
	less  func(a, b T) bool // true 表示 a 的优先级更高
}

// heap.Interface 是 container/heap 要求堆的底层数据结构要实现的适配接口；
// 其中 Push / Pop 负责底层切片的追加和移除，它们不会维护堆序。
// 通过 heap.Push、heap.Pop 函数操作底层数据，才会维护堆序，形成优先级队列行为。
// 编译期检查：PriorityQueue[T] 实现了 heap.Interface
var _ heap.Interface = (*HeapList[any])(nil)

func NewHeapList[T any](less func(a, b T) bool) *HeapList[T] {
	return &HeapList[T]{less: less}
}

func (pq HeapList[T]) Len() int           { return len(pq.items) }
func (pq HeapList[T]) Less(i, j int) bool { return pq.less(pq.items[i], pq.items[j]) }
func (pq HeapList[T]) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }

// 以下两个方法是 heap.Interface 要求的 any 签名

func (pq *HeapList[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}

func (pq *HeapList[T]) Pop() any {
	n := len(pq.items)
	item := pq.items[n-1]
	var zero T
	pq.items[n-1] = zero // 避免底层数组继续引用该元素
	pq.items = pq.items[:n-1]
	return item
}

func (pq *HeapList[T]) Peek() T {
	if pq.Len() == 0 {
		panic(fmt.Errorf("empty heap"))
	}
	return pq.items[0]
}
