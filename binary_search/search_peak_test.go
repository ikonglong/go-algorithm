package binary_search

// https://leetcode.cn/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)-1
	return searchPeak(arr, low, high)
}

// 采用小争哥的二分代码模板
// 注意，这个实现不能处理首元素或尾元素就是峰值的输入
func searchPeak(arr []int, low, high int) int {
	n := len(arr)
	for low <= high {
		mid := low + (high-low)/2
		// 由下面的 mid-1>=0 => mid>=1 => mid>0 => case: mid=0
		// 由下面的 mid+1<n => mid<n-1 => case: mid=n-1
		if mid == 0 {
			// 可以这么做是因为题目说不会给首元素就是峰值的输入
			low = mid + 1
		} else if mid == n-1 {
			// 可以这么做是因为题目说不会给尾元素就是峰值的输入
			high = mid - 1
		} else if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] < arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 实现OK，但没有采用小争哥给的二分代码模板
// 注意，这个实现可以处理首元素或尾元素就是峰值的输入
func searchPeak2(arr []int, low, high int) int {
	if low == high {
		return low
	}
	mid := (low + high) / 2
	// 根据 mid 相邻元素的变化趋势判断目标元素在哪边
	// 这里判断时为什么不采用 arr[mid-1] < arr[mid]？
	// 因为计算 mid 时只要 high>low 就向下取整，这意味着mid+1一定不会越界
	if arr[mid] < arr[mid+1] {
		return searchPeak2(arr, mid+1, high)
	} else if arr[mid] > arr[mid+1] {
		return searchPeak2(arr, low, mid)
	} else {
		// 根据题意，没有相等的元素，因此不可能走到这里
		return -1
	}
}
