package twopointer

import "slices"

// https://leetcode.cn/problems/3sum/

// 既然是求满足条件的数对，那么首要问题，或者说最基本的问题是：怎么划分数对？

// 算法思路：
// 遍历数组，把当前元素作为第一个数 a，然后在数组的剩余部分找和为 target-a 的所有数对。
// 这样就把三数之和分解为遍历数组确定第一个数 a 和两数之和问题了。
//
// 如何判断答案是否重复？
//  1. 添加答案时，遍历答案元素列表，逐个比对是否相同
//  2. 将每个答案转换为字符串，用它作为答案的唯一标识，放入哈希表用来判重
//  3. 第 1 层循环确定的是第 1 个数，因此在该层循环的开头判断当前元素是否跟前一个相等，
//     若相等则跳过后续逻辑，进行下次迭代，这样就实现了对第 1 个数的去重。
//     第 2 层循环确定的是后两个数，因此在这层循环的开头对第 2 个数进行去重，去重逻辑类似。
//     对前两个数去重后，实际上不需要对第三个数去重了，因为只要前两个数不同，答案肯定不重复
//
// 为什么计算两数之和也需要去重？
// 因为输入 [0,0,0,0]，取第一个作为答案的第一个数，在剩余部分上计算两数之和依然会产出重复解

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	n := len(nums)
	slices.Sort(nums)
	result := make([][]int, 0)

	for i, a := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 对答案的第一个数去重
			continue
		}
		low, high := i+1, n-1
		twoSum := -a
		for low < high {
			if low > i+1 && nums[low] == nums[low-1] { // 对答案的第二个数去重
				low++
				continue
			}
			// 不需要对答案的第三个数去重
			// if high < n-1 && nums[high] == nums[high+1] { // 对答案的第三个数去重
			// 	high--
			// 	continue
			// }
			s := nums[low] + nums[high]
			if s == twoSum {
				result = append(result, []int{a, nums[low], nums[high]})
				// low++ or high-- 也可。但更推荐如下写法，因为每个元素只能用一次
				low++
				high--
			} else if s < twoSum {
				low++
			} else {
				high--
			}
		}
	}

	return result
}

// 此实现版本将两数之和问题的求解提取为单独的函数
func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	slices.Sort(nums)
	result := make([][]int, 0)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] { // 对答案的第一个数去重
			continue
		}
		solutions := twoSum(nums, i+1, -a)
		for _, pair := range solutions {
			result = append(result, append([]int{a}, pair...))
		}
	}
	return result
}

// 双指针实现思路
func twoSum(nums []int, start, target int) [][]int {
	result := make([][]int, 0)
	low, high := start, len(nums)-1
	for low < high {
		if low > start && nums[low] == nums[low-1] { // 对答案的第二个数去重
			low++
			continue
		}
		// 不需要对答案的第三个数去重
		// if high < len(nums)-1 && nums[high] == nums[high+1] { // 对答案的第三个数去重
		// 	high--
		// 	continue
		// }
		s := nums[low] + nums[high]
		if s == target {
			result = append(result, []int{nums[low], nums[high]})
			// low++ or high-- 也可。但更推荐如下写法，因为每个元素只能用一次
			low++
			high--
		} else if s < target {
			low++
		} else {
			high--
		}
	}
	return result
}
