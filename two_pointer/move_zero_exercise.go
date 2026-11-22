package twopointer

// https://leetcode.cn/problems/move-zeroes/

func moveZeroes_(nums []int) {
	p, q := 0, 0
	for q < len(nums) {
		if nums[q] == 0 {
			q++
		} else {
			nums[p], nums[q] = nums[q], nums[p]
			p++
			q++
		}
	}
}

func moveZeroes2_(nums []int) {
	p := 0
	for i :=0; i<len(nums);i++{
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
}
