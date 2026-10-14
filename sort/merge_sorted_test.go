package sort_

// https://leetcode.cn/problems/sorted-merge-lcci/description/
// 合并两个有序数组 a、b，a 的末端有足够的空间容纳 b

// 插入排序，将 B 逐个插入 A 中
func merge(A []int, m int, B []int, n int) {
	for j := range n {
		v := B[j]
		i := m - 1
		for ; i >= 0 && A[i] > v; i-- {
			// 遍历过程中，把每个元素存入后一格，腾出插入位置
			A[i+1] = A[i]
		}
		// 退出循环时就找到了插入位置
		if i < 0 { // 所有元素都大于 v
			A[0] = v
		} else { // A[i] <= v
			A[i+1] = v
		}
		// 添加一个元素后，m 自增
		m++
	}
}

// 类似归并排序，A 和 B 两个数组都已有序，将它们看作两个队列。
// 创建一个新数组，表示已排序的结果数组。创建两个指针分别指向
// 两个队列的头元素，比较头元素较小的出队。任何一个队列为空了，
// 就将另一个队列剩余部分直接放入结果数组
func merge2(A []int, m int, B []int, n int) {
	arr := make([]int, 0, m+n)
	pa, pb := 0, 0
	for pa < m && pb < n {
		va, vb := A[pa], B[pb]
		if va < vb {
			arr = append(arr, va)
			pa++
		} else {
			arr = append(arr, vb)
			pb++
		}
	}
	que := A
	start := pa
	end := m
	// 谁先遍历完跟队列大小没关系，而是由元素间的相对大小决定的。
	// 例如，A 元素较多，而 B 只有一个元素，但 A 中的所有元素都小于 B，那么 A 也会先遍历完
	//if m > n {
	if pb < n {
		que = B
		start = pb
		end = n
	}
	for i := start; i < end; i++ {
		arr = append(arr, que[i])
	}
	for i, v := range arr {
		A[i] = v
	}
}

// 改进上一个归并思路实现版本：不在循环后再检查哪个队列有剩余并继续处理，而是直接在循环中完成处理
func merge3(A []int, m int, B []int, n int) {
	// arr := make([]int, 0, m+n)
	arr := make([]int, m+n)
	pa, pb := 0, 0
	for pa < m || pb < n {
		v := 0
		if pa == m {
			v = B[pb]
			pb++ // 别忘了更新循环变量
		} else if pb == n {
			v = A[pa]
			pa++
		} else if A[pa] < B[pb] {
			v = A[pa]
			pa++
		} else {
			v = B[pb]
			pb++
		}
		// arr = append(arr, v)
		// pa, pb 之前的元素都已添加至结果中，现在要将其中一个位置的值添加至结果中，
		// 剩余一个还未处理，因此减 1
		arr[pa+pb-1] = v
	}
	for i, v := range arr {
		A[i] = v
	}
}

// 从左到右归并两个数组，必须创建一个新的结果数组，无法直接利用 A 后端的缓冲区。
// 因为 A 后端有缓冲区，因此我们可以
// 1. 先计算出合并后数组的长度，即确定最后一个元素会写入的位置 m+n-1
// 2. 创建一个指针 tail 指向合并后数组尾部的可写入位置，初始时为 m+n-1
// 3. 创建两个指针 pa, pb，分别指向两个数组的尾元素。比较尾元素，较大的出队尾，并移动对应的指针
func merge4(A []int, m int, B []int, n int) {
	at, bt := m-1, n-1
	tail := m + n - 1
	for at >= 0 || bt >= 0 {
		v := 0
		// 因为 at、bt 为 0 还要迭代，因此等于 0 不是处理完的标志，等于 -1 才是
		// A:[1,2,3,0,0,0],B:[2,5,6]，at=0时，bt=-1，逻辑正确时走第二个分支，
		// 但却走了第一个分支，导致数组越界
		//if at == 0 {
		if at == -1 {
			v = B[bt]
			bt--
		} else if bt == -1 {
			v = A[at]
			at--
		} else if A[at] > B[bt] {
			v = A[at]
			at--
		} else {
			v = B[bt]
			bt--
		}
		A[tail] = v
		tail--
	}
}
