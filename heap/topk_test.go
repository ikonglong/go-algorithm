package heap

import "container/heap"

// https://leetcode.cn/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	freqMap := make(map[int]int)
	for _, e := range nums {
		freqMap[e] = freqMap[e] + 1
	}

	heapList := NewHeapList(func(a, b QElement2) bool {
		// bug：这是大顶堆。求 Top K 大的，应该用小顶堆
		// return a.count > b.count
		return a.count < b.count
	})
	heap.Init(heapList)

	for val, count := range freqMap {
		item := QElement2{
			val:   val,
			count: count,
		}
		// bug：应该用堆大小比较
		// if len(freqMap) < k {
		if heapList.Len() < k {
			heap.Push(heapList, item)
		} else if heapList.Peek().count < count {
			heap.Pop(heapList)
			heap.Push(heapList, item)
		}
	}

	result := make([]int, 0)
	for heapList.Len() > 0 {
		result = append(result, heap.Pop(heapList).(QElement2).val)
	}
	return result
}

type QElement2 struct {
	val   int
	count int
}
