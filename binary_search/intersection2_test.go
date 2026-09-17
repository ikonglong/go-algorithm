package binary_search

import (
	"fmt"
	"slices"
	"testing"
)

func intersect2(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	small, big := nums1, nums2
	if len(nums1) > len(nums2) {
		small, big = nums2, nums1
	}

	rs := make([]int, 0, len(small))
	for i := 0; i < len(small); i++ {
		if len(rs) == 0 || rs[len(rs)-1] != small[i] {
			if present, cnt := find2(big, small[i]); present {
				for j := i; j < len(small) && small[j] == small[i] && j < i+cnt; j++ {
					rs = append(rs, small[j])
				}
			}
		}
	}
	return rs
}

func find2(nums []int, target int) (bool, int) {
	start, end, mid := 0, len(nums)-1, 0
	for {
		if start == end {
			if target == nums[start] {
				cnt := 1
				for i := start + 1; i < len(nums) && nums[i] == target; i++ {
					cnt++
				}
				for i := start - 1; i >= 0 && nums[i] == target; i-- {
					cnt++
				}
				fmt.Printf("nums:%v, target: %v, present: %v, cnt: %v\n", nums, target, true, cnt)
				return true, cnt
			}
			fmt.Printf("nums:%v, target: %v, present: %v, cnt: %v\n", nums, target, false, 0)
			return false, 0
		}

		mid = (start + end) / 2
		if target <= nums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
}

func TestIntersection2(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tc := range testCases {
		got := intersect2(tc.nums1, tc.nums2)
		if !slices.Equal(got, tc.want) {
			t.Errorf("intersect2(%v, %v) = %v, want: %v",
				tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
