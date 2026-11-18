package heap

import "container/heap"

// https://leetcode.cn/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	heapList := NewHeapList[QElement](func(a, b QElement) bool {
		// true 表示 a 的优先级更高
		return a.currNode.Val < b.currNode.Val
	})
	heap.Init(heapList)

	// 将 k 个链表的头节点放入优先级队列 min-heap
	for _, head := range lists {
		if head == nil {
			continue
		}
		heap.Push(heapList, QElement{
			currNode: head,
		})
	}

	dummyHead := &ListNode{}
	tail := dummyHead
	for heapList.Len() > 0 {
		topNode := heap.Pop(heapList).(QElement).currNode

		tail.Next = topNode
		tail = topNode

		// 注意，这里是合并多个有序链表，不是求 TopK，因此不存在比较下一个元素和堆顶，
		// 然后决定是否删除堆顶元素的逻辑。每次迭代总是先弹出堆顶元素，添加到链表尾，再
		// 将堆顶节点的下一个节点补充到堆中。这样一直维持堆大小为 k，直到有链表遍历完了
		if topNode.Next != nil {
			heap.Push(heapList, QElement{currNode: topNode.Next})
		}
	}

	return dummyHead.Next
}

type QElement struct {
	currNode *ListNode
}
