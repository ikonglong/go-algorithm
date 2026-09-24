package twopointer

import "slices"

// https://leetcode.cn/problems/two-sum/
// 题目说，每个数只有一个答案

// 在给定数组中寻找和为 target 的数对。有两种方法分解问题：
// 1. 既然是两数 a、b 之和，那么另一个数就是 target-a
// 2. 先用某种方式将数组划分成一个个数对；再计算和，跟 target 比较

// 修改一下题目细节，提高一点难度：
// 假设解的某个元素在数组中存在重复，求不重复的解。
// 例如：nums=[1,2,4,2], target=6, 输出：
//
// 双指针解法
func twoSumVariant3(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	// 因为双指针解法需要数组有序，但结果要求返回下标，因此不能直接对原数组排序
	sortedNums := append([]int(nil), nums...)
	slices.Sort(sortedNums)
	// 为了去除重复记录，元素是否已使用。找到解时，需在其中查找第一个已使用的元素
	used := make([]bool, n)

	low, high := 0, n-1
	for low < high {
		sum := sortedNums[low] + sortedNums[high]
		if sum > target {
			high--
		} else if sum < target {
			low++
		} else {
			lowIdx, isLowUsed := checkUsed(sortedNums[low], nums, used)
			if isLowUsed { // 跟之前重复
				continue
			}
			used[lowIdx] = true

			highIdx, isHighUsed := checkUsed(sortedNums[high], nums, used)
			if isHighUsed {
				continue
			}
			used[highIdx] = true

			return []int{lowIdx, highIdx}
		}
	}
	return nil
}

func checkUsed(value int, nums []int, used []bool) (int, bool) {
	for i, v := range nums {
		if v == value {
			return i, used[i]
		}
	}
	return -1, false
}
