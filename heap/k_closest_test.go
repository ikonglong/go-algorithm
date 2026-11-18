package heap

import "container/heap"

// https://leetcode.cn/problems/k-closest-points-to-origin/
// 求 K 个距离原点最近的点

func kClosest(points [][]int, k int) [][]int {
	// 大顶堆
	maxHeap := NewHeapList(func(a, b QElement3) bool {
		return a.dist > b.dist
	})

	i := 0
	for ; i < k && i < len(points); i++ {
		heap.Push(maxHeap, QElement3{
			dist:  points[i][0]*points[i][0] + points[i][1]*points[i][1],
			index: i,
		})
	}

	for ; i < len(points); i++ {
		dist := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if dist < maxHeap.Peek().dist {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, QElement3{
				dist:  dist,
				index: i,
			})
		}
	}

	ans := make([][]int, 0)
	for maxHeap.Len() > 0 {
		item := heap.Pop(maxHeap).(QElement3)
		ans = append(ans, points[item.index])
	}
	return ans
}

type QElement3 struct {
	dist  int
	index int
}
