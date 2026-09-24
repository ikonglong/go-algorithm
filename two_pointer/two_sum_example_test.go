package twopointer

import "slices"

// 头尾指针
func pairSums(nums []int, target int) [][]int {
	result := make([][]int, 0)

	slices.Sort(nums)
	// low、high 分别指向数组头、尾
	// low 只能单调递增，high 只能单调递减
	low, high := 0, len(nums)-1

	// test case: -10 -7 5 12 15
	// 两个指针相遇也应该结束
	for low < high {
		sum := nums[low] + nums[high]
		if sum > target {
			high--
		} else if sum < target {
			low++
		} else {
			result = append(result, []int{nums[low], nums[high]})
			high--
			low++
		}
	}
	return result
}
