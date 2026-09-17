package binary_search

import (
	"fmt"
	"strings"
	"testing"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end, mid, depth := 0, len(nums)-1, 0, 0
	var indent, indent2 string
	for {
		indent = strings.Repeat(" ", 4*depth)
		indent2 = strings.Repeat(" ", 4*(depth+1))
		fmt.Printf("%sstart:%d, end:%d, nums:%v\n",
			indent, start, end, nums[start:end+1])

		if start == end {
			if target == nums[start] {
				return start
			}
			return -1
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
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, tc := range testCases {
		got := search(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
