package binary_search

// https://leetcode.cn/problems/find-peak-element/

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	low, high := 0, len(nums)-1
	return searchPeakX(nums, low, high)
}

// 实现OK，但没有采用小争哥给的二分代码模板
// 注意，这个实现可以处理首元素或尾元素就是峰值的输入
func searchPeakX(arr []int, low, high int) int {
	if low == high {
		return low
	}
	mid := (low + high) / 2
	// 根据 mid 相邻元素的变化趋势判断目标元素在哪边
	// 这里判断时为什么不采用 arr[mid-1] < arr[mid]？
	// 因为计算 mid 时只要 high>low 就向下取整，这意味着mid+1一定不会越界
	if arr[mid] < arr[mid+1] {
		return searchPeakX(arr, mid+1, high)
	} else if arr[mid] > arr[mid+1] {
		return searchPeakX(arr, low, mid)
	} else {
		// 根据题意，没有相等的元素，因此不可能走到这里
		return -1
	}
}

// 采用循环实现前一个版本。OK
func searchPeakX2(a []int, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if lo == hi {
			return lo
		}
		if a[mid] < a[mid+1] { // mid 在上坡区间
			lo = mid + 1
		} else { // a[mid]>=a[mid+1] => mid 在下坡区间
			hi = mid
		}
	}
	return -1
}

// Not OK，无法处理首元素或尾元素就是峰值的情形
func searchPeakX3(a []int, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if lo == hi {
			return lo
		}
		if a[mid-1] < a[mid] { // mid 在上坡区间
			// 对于用例[1,2,3,1]，输出 3，期望 2
			lo = mid + 1
			// 对于用例[1,2,3,1]，超出时间限制，应该是无限循环了
			// lo = mid
		} else { // a[mid-1]>=a[mid] => mid 在下坡区间
			hi = mid - 1
		}
	}
	return -1
}

// 参考小争哥的习题讲解写的版本。当然肯定采用了小真哥的二分代码模板
func findPeakElement4(a []int) int {
	if len(a) == 1 {
		return 0
	}
	n := len(a)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// if lo == hi { // 处理输入为[1]，a[mid+1]越界情形。也可以在方法开头处理
		// 	return mid
		// }
		// 下面期望 mid-1>=0，mid+1<n
		// mid-1>=0 => mid>=1 => mid>0 => 想到看看 mid=0 是啥情况
		// mid+1<n => mid<n-1 => 想看看 mid=n-1 是啥情况
		if mid == 0 { // 首元素是峰值的情形
			if a[mid] > a[mid+1] {
				return mid
			} else {
				lo = mid + 1
			}
		} else if mid == n-1 { // 尾元素是峰值的情形
			if a[mid-1] < a[mid] {
				return mid
			} else {
				hi = mid - 1
			}
		} else if a[mid-1] < a[mid] && a[mid] > a[mid+1] {
			return mid
		} else if a[mid] < a[mid+1] {
			// 题目说可能有多个峰值，这意味着上坡的右侧一定有一个峰值，
			// 而上坡的左侧可能有峰值，也可能没有
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// 想套小争哥的代码模板，但踩坑了
//
// 对于输入：[1]
// 这样写：
//
//	if mid == 0 {
//		lo = mid
//	} else if mid == n-1 {
//		hi = mid
//	}
//
// 对于用例[2,1]，超出时间限制，应该是无限循环了
//
// 这样写：
//
//	if mid == 0 {
//		lo = mid+1
//	} else if mid == n-1 {
//		hi = mid-1
//	}
//
// 对于用例[2,1]，输出：-1，期望 0
func searchPeakX5(a []int, lo, hi int) int {
	n := len(a)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// 由下面的 mid-1>=0 => mid>=1 => mid>0 => case: mid=0
		// 由下面的 mid+1<n => mid<n-1 => case: mid=n-1
		if mid == 0 {
			lo = mid + 1
		} else if mid == n-1 {
			hi = mid - 1
		} else if a[mid-1] < a[mid] && a[mid] > a[mid+1] {
			return mid
		} else if a[mid-1] < a[mid] { // mid 在上坡区间
			lo = mid + 1
		} else { // mid 在下坡区间
			hi = mid - 1
		}
	}
	return -1
}
