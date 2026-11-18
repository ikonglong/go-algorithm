package heap

import "container/heap"

// https://leetcode.cn/problems/find-median-from-data-stream/

type MedianFinder struct {
	minHeap *HeapList[int]
	maxHeap *HeapList[int]
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewHeapList(func(a, b int) bool {
			return a < b
		}),
		maxHeap: NewHeapList(func(a, b int) bool {
			return a > b
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 初始时两个堆都为空，新数据无法跟堆顶元素比较大小，此时直接放入大顶堆
	if this.maxHeap.Len() == 0 || num <= this.maxHeap.Peek() {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	// 每次插入新数据后，通过在两个堆之间移动堆顶元素以维护两个堆的大小比例接近 50:50
	// 元素总数 n 为奇数时，大顶堆应该比小顶堆大 1
	for this.maxHeap.Len() < this.minHeap.Len() {
		val := heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, val)
	}
	for this.minHeap.Len() < this.maxHeap.Len()-1 {
		val := heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// 若大顶堆大小大于小顶堆，则表明已处理的数据个数为奇数
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Peek())
	} else { // 若大小相等，表明已处理的数据个数为偶数，中位数就取中间两个数的平均值
		return float64(this.maxHeap.Peek()+this.minHeap.Peek()) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
