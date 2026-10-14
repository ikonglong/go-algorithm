package binary_search

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 真解的值是确定的

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	return []int{searchByPos(nums, lo, hi, target, true), searchByPos(nums, lo, hi, target, false)}
	// return []int{searchByPos2(nums, target, true), searchByPos2(nums, target, false)}
}

// 递归实现
func searchByPos(nums []int, lo, hi, target int, firstOrLast bool) int {
	// 最小子问题？
	if lo > hi { // 这个检测等价于循环实现版本中的循环退出条件
		return -1
	}
	// 下面的逻辑等价于循环实现版本中的循环体逻辑
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		if firstOrLast {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				hi = mid - 1
				return searchByPos(nums, lo, hi, target, firstOrLast)
			}
		} else {
			if mid == hi || nums[mid+1] != target {
				return mid
			} else {
				lo = mid + 1
				return searchByPos(nums, lo, hi, target, firstOrLast)
			}
		}
	} else if nums[mid] < target {
		lo = mid + 1
		return searchByPos(nums, lo, hi, target, firstOrLast)
	} else {
		hi = mid - 1
		return searchByPos(nums, lo, hi, target, firstOrLast)
	}
	// return -1
}

// 循环实现
func searchByPos2(nums []int, target int, firstOrLast bool) int {
	n := len(nums)
	lo, hi := 0, n-1
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			if firstOrLast { // search first
				// 若前一个元素不等于 target，那么 nums[mid] 就是第一个
				// mid-1>=0 => mid>=1 => mid>0 => 想到看看 mid == 0 是什么情况
				if mid == 0 || nums[mid-1] != target {
					result = mid
				} else {
					hi = mid - 1
				}
			} else { // search last
				// 若后一个元素不等于 target，那么 nums[mid] 就是最后一个
				// mid+1<n => mid<n-1 => 想到看看 mid == n-1 是什么情况
				if mid == n-1 || nums[mid+1] != target {
					result = mid
				} else {
					lo = mid + 1
				}
			}
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}
