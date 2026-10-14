package sort_

// https://leetcode.cn/problems/smallest-k-lcci/
// 返回最小的k个数

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	lo, hi := 0, len(arr)-1
	kthIdx := smallestKth(arr, lo, hi, k)
	if kthIdx == -1 {
		return nil
	}
	return arr[:kthIdx+1]
}

func smallestKth(a []int, lo, hi, k int) int {
	p := partition_asc(a, lo, hi)
	if k == p-lo+1 {
		return p
	} else if p-lo+1 < k {
		return smallestKth(a, p+1, hi, k-(p-lo+1))
	} else {
		return smallestKth(a, lo, p-1, k)
	}
}

func partition_asc(a []int, lo, hi int) int {
	if lo == hi { // 避免对空子数组进行分割
		return lo
		// 埋有 bug
		// return -1
	}

	i, j := lo, hi+1
	// p(ivot) 枢轴元素
	p := a[lo]
	for {
		// 左扫描，跳过小于枢轴的，停在>=枢轴的或者hi
		for i++; i < hi && a[i] < p; i++ {
		}
		// 右扫描，跳过大于枢轴的，停在<=枢轴的或者lo
		for j--; j > lo && p < a[j]; j-- {
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
