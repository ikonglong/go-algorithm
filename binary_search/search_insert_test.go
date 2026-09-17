package binary_search

import (
	"fmt"
	"strings"
	"testing"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return searchInRange(nums, target, 0, len(nums)-1, 0)
}

func searchInRange(nums []int, target int, start, end, depth int) int {
	indent := strings.Repeat(" ", 4*depth)
	fmt.Printf("%sstart:%d, end:%d, nums:%v\n", indent, start, end, nums[start:end+1])
	if end == start {
		if target <= nums[start] {
			return start
		} else {
			return start + 1
		}
	}
	mid := start + (end-start)/2
	indent2 := strings.Repeat(" ", 4*(depth+1))
	fmt.Printf("%smid:%d, nums[mid]:%d, target:%d\n", indent2, mid, nums[mid], target)
	if target <= nums[mid] {
		fmt.Printf("%sleft half, nums[%d-%d]:%v\n",
			indent2, start, mid, nums[start:mid+1])
		// mid+1 导致无限循环
		// return searchInRange(nums, target, start, mid+1)
		return searchInRange(nums, target, start, mid, depth+1)
	} else {
		fmt.Printf("%sright half, nums[%d-%d]:%v\n",
			indent2, mid+1, end, nums[mid+1:end+1])
		return searchInRange(nums, target, mid+1, end, depth+1)
	}
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

func searchInsertV2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end, mid, depth := 0, len(nums)-1, 0, 0
	var indent, indent2 string
	for {
		indent = strings.Repeat(" ", 4*depth)
		indent2 = strings.Repeat(" ", 4*(depth+1))
		fmt.Printf("%sstart:%d, end:%d, nums:%v\n",
			indent, start, end, nums[start:end+1])

		if start == end {
			if target <= nums[start] {
				return start
			}
			return start + 1
		}

		mid = (start + end) / 2
		fmt.Printf("%smid:%d, nums[mid]:%d, target:%d\n",
			indent2, mid, nums[mid], target)
		if target <= nums[mid] {
			end = mid
			fmt.Printf("%sleft half, nums[%d-%d]:%v\n",
				indent2, start, mid, nums[start:mid+1])
		} else {
			start = mid + 1
			fmt.Printf("%sright half, nums[%d-%d]:%v\n",
				indent2, mid+1, end, nums[mid+1:end+1])
		}
		depth += 1
	}
}

func TestSearchInsertV2(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		// {[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		// {[]int{1, 3, 5, 6}, 7, 4},
		// {[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tc := range testCases {
		got := searchInsertV2(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("searchInsert(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
