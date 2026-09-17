package binary_search

import (
	"fmt"
	"slices"
	"testing"
)

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	small, big := nums1, nums2
	if len(nums2) < len(nums1) {
		small = nums2
		big = nums1
	}
	rs := make([]int, 0, len(small))
	fmt.Printf("small:%v, big:%v\n", small, big)

	for _, item := range small {
		if len(rs) == 0 || (len(rs) > 0 && rs[len(rs)-1] != item) { // 去重
			if find(big, item) {
				rs = append(rs, item)
			}
			fmt.Printf("item: %d, rs: %v\n", item, rs)
		}
	}
	return rs
}

func find(nums []int, target int) bool {
	start, end, mid := 0, len(nums)-1, 0
	for {
		if start == end {
			present := target == nums[start]
			fmt.Printf("nums: %v, target: %d, present:%v\n",
				nums, target, present)
			return present
		}

		mid = (start + end) / 2
		if target <= nums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tc := range testCases {
		got := intersection(tc.nums1, tc.nums2)
		if !slices.Equal(got, tc.want) {
			t.Errorf("intersection(%v, %v) = %v, want: %v",
				tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
