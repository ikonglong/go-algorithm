package twopointer

import "slices"

// https://leetcode.cn/problems/3sum/description/

// 既然是求满足条件的数对，那么先将问题分解为：
// 1. 怎么划分数对？
// 2. 判断数对是否满足要求

// 算法思路：
// 求和为零的三个元素：先从数组中选择一个元素作为其中之一，再计算另两个元素的和是多少，
// 这样就把三数之和分解为数组遍历、计算两数之和
//
// 如何判断可行解是否重复？
//  1. 添加新解时，遍历解列表逐个比对是否相同
//  2. 将每个可行解转换为字符串，用它作为这个可行解的唯一标识，放入哈希表。
//  3. 外层循环中，若当前元素跟前一个相等，则跳过后续逻辑，直接下一次迭代；
//     计算两数之和时，若 low 跟前一个相等，或者 high 跟后一个相等，也需跳过，直接下一次迭代。
//     为什么计算两数之和也需要去重？例如：[0,0,0,0]，取第一个作为解的第一个，剩余计算两数之和
//     依然会产出重复解

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	n := len(nums)
	slices.Sort(nums)
	result := make([][]int, 0)

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 去除解中第一个元素的重复
			continue
		}
		low, high := i+1, n-1
		twoSum := 0 - v
		for low < high {
			if low > i+1 && nums[low] == nums[low-1] { // 去除解中第二个元素的重复
				low++
				continue
			}
			if high < n-1 && nums[high] == nums[high+1] { // 去除解中第三个元素的重复
				high--
				continue
			}
			sum := nums[low] + nums[high]
			if sum == twoSum {
				result = append(result, []int{v, nums[low], nums[high]})
				high-- // or low++ or low++&&high--
			} else if sum < twoSum {
				low++
			} else {
				high--
			}
		}
	}

	return result
}

// 将两数之和提取为函数的实现版本
func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	slices.Sort(nums)

	result := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && v == nums[i-1] { // 去除第一个元素的重复
			continue
		}
		solutions := twoSum(nums, i+1, 0-v)
		for _, list := range solutions {
			result = append(result, append([]int{v}, list...))
		}
	}
	return result
}

//  1. 先划分出数对：创建两个指针，low、high，分别指向 0、n-1
//  2. 若数对和大于 target，high--；若小于，low++；
//     若等于，找到一个可行解，然后 h-- 或 l++。
//     注意，low high 指针不能往回走，否则就会产出重复解
//     重复上述步骤，直到 low 大于等于 high

func twoSum(nums []int, start, target int) [][]int {
	result := make([][]int, 0)
	low, high := start, len(nums)-1
	for low < high {
		if low > start && nums[low] == nums[low-1] { // 去除第二个元素的重复
			low++
			continue
		}
		if high < len(nums)-1 && nums[high] == nums[high+1] { // 去除第三个元素的重复
			high--
			continue
		}
		sum := nums[low] + nums[high]
		if sum == target {
			result = append(result, []int{nums[low], nums[high]})
			high-- // or low++ or high-- && low++
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return result
}
