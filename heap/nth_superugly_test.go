package heap

import "container/heap"

// https://leetcode.cn/problems/super-ugly-number/

func nthSuperUglyNumber(n int, primes []int) int {
	minHeap := NewHeapList(func(a, b int64) bool {
		// bug：这是大顶堆，应该用小顶堆
		// return a > b
		return a < b
	})
	numSet := make(map[int64]struct{})

	// 题目说 1 是超级丑数，不含质因数，即它的所有质因数都在给定的质数数组中
	heap.Push(minHeap, int64(1))
	numSet[int64(1)] = struct{}{}

	// 注意，这里的实现思路【不是】维持堆大小为 n，不断插入扩展出来的丑数，最后取顶堆元素，
	// 【而是】重复如下操作：
	// a. 堆顶元素出堆
	// b. 扩展堆顶元素，将产生的丑数插入堆
	// 第 n 次出堆的堆顶元素就是期望结果

	count := 0 // 出堆次数
	num := int64(1)
	for minHeap.Len() > 0 && count < n {
		// 堆顶元素出队，扩展它，产生下一层丑数
		num = heap.Pop(minHeap).(int64)
		count++

		// 答案已经拿到，可以直接返回，无需再扩展
		if count == n {
			return int(num)
		}

		for _, p := range primes {
			uglyNum := num * int64(p)
			if _, found := numSet[uglyNum]; !found {
				numSet[uglyNum] = struct{}{}
				heap.Push(minHeap, uglyNum)
			}
		}
	}

	return int(num)
}
