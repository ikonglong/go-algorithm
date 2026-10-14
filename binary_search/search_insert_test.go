package binary_search

import (
	"testing"
)

// https://leetcode.cn/problems/search-insert-position/description/
// 搜索插入位置

// 递归实现的二分搜索
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	return doSearchInsert(nums, low, high, target)
}

// 返回插入位置
func doSearchInsert(nums []int, low, high int, target int) int {
	// 若 target 大于所有元素，low 会移动至 n；
	// 若 target 小于所有元素，high 会移动至 -1，low 为 0
	// 若 target 未找到，那么先想象找到时候的情形，low、high 都指向它，未找到时 low、high
	// 为了找它而错过彼此，此时 low 就指向下一个元素，正好也是 target 的插入位置
	// 不管哪种情形，按顺序插入 targets 的位置都为 low
	if low > high {
		return low
	}
	mid := (low + high) / 2
	if target == nums[mid] {
		return mid
	} else if nums[mid] < target {
		return doSearchInsert(nums, mid+1, high, target)
	} else {
		return doSearchInsert(nums, low, mid-1, target)
	}
}

// 循环实现的二分搜索
func searchInsertV2(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	// 若 target 大于所有元素，low 会移动至 n；
	// 若 target 小于所有元素，high 会移动至 -1，low 为 0
	// 若 target 未找到，那么先想象找到时候的情形，low、high 都指向它，未找到时 low、high
	// 为了找它而错过彼此，此时 low 就指向下一个元素，正好也是 target 的插入位置
	// 不管哪种情形，按顺序插入 targets 的位置都为 low
	return low
}

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tc := range testCases {
		got := searchInsert(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("searchInsert(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}

func searchInsert3(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return doSearchInsert3(nums, 0, len(nums)-1, target)
}

// 返回 target 所在位置，若 target 未找到，则返回其插入位置
func doSearchInsert3(nums []int, low, high, target int) int {
	if low == high {
		// 若找到 target，low 就是其下标；若未找到，nums[low/high] 也是最接近它的元素。
		// 这里的思路是：找到最接近它的元素时，计算其插入位置，避免因 low 和 high 错过
		// 致使 low 大于 high
		if target <= nums[low] {
			return low
		} else {
			return low + 1
		}
	}
	mid := low + (high-low)/2
	if target <= nums[mid] {
		// mid+1 导致无限循环
		// return doSearchInsert3(nums, low, mid+1, target)
		return doSearchInsert3(nums, low, mid, target)
	} else {
		return doSearchInsert3(nums, mid+1, high, target)
	}
}
