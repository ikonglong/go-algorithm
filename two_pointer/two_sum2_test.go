package twopointer

import (
	"slices"
)

// https://leetcode.cn/problems/two-sum/
// 题目说，每个数只有一个答案

// 在给定数组中寻找和为 target 的数对。有两种方法分解问题：
// 1. 既然是两数 a、b 之和，那么另一个数就是 target-a
// 2. 先用某种方式将数组划分成一个个数对；再计算和，跟 target 比较

// 双指针解法
func twoSumVariant2(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	// 因为双指针解法需要数组有序，但结果要求返回下标，因此不能直接对原数组排序
	sortedNums := append([]int(nil), nums...)
	slices.Sort(sortedNums)
	// 针对同一个解中两个元素相同的 case，例如：输入 [3,3]
	// 记录元素在解中是否被使用过
	used := make([]bool, n)

	low, high := 0, n-1
	for low < high {
		sum := sortedNums[low] + sortedNums[high]
		if sum > target {
			high--
		} else if sum < target {
			low++
		} else {
			// 我的脑子真TM蠢。反复犯同一个错：originalIdx(nums[low], ...)
			origLow := find(sortedNums[low], nums, used)
			origHigh := find(sortedNums[high], nums, used)
			return []int{origLow, origHigh}
		}
	}
	return nil
}

func find(value int, nums []int, used []bool) int {
	i := 0
	for ; i < len(nums); i++ {
		// 如果有重复元素，但已使用的，则跳过；
		// 找到第一个未使用的，将其标记为已使用，返回其下标
		if nums[i] == value && used[i] == false {
			used[i] = true
			// return i
			break
		}
	}
	// panic(fmt.Errorf("not found unused element(%d)", value))
	return i
}

// 先确定其中一个数，再求出另一个数，然后在哈希表中查找它。
// 跟 v3 相比，这个版本效率略低。因为它是先构建索引，再遍历数组查找另一个数；
// 而 v3 是在遍历数组的过程中，边查找边构建索引，如果找到就直接返回了。
func twoSumVariant2_v2(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	numMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numMap[v] = i
	}

	// 有重复值也能工作。虽然在 numMap 中后面覆盖前面，
	// 但遍历数组时仍可访问前面的元素，如果另一个数刚好是后面的重复，
	// 结果就是期望的。例如：[3,3]

	for i, x := range nums {
		// 想通过删除自己来避免自己匹配自己。但若 v1 跟后面的某个元素相同，
		// 就会把后面的元素一并删除了。例如：[3,3]
		// delete(numMap, x)

		y := target - x
		// 通过 i != j 避免自己跟自己匹配
		if j, found := numMap[y]; found && i != j {
			return []int{i, j}
		}
	}
	return nil
}

// 先确定其中一个数，再求出另一个数，然后在哈希表中查找它
func twoSumVariant2_v3(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	numMap := make(map[int]int)

	for i, x := range nums {
		// 心智模型是：把 Map 想成一本“过去出现过的数的登记簿”。每个新来的数，
		// 先查有没有自己的搭档；没有，就登记自己，等后面的数来找。
		// 「先查、后存」这个顺序的目的是避免自己匹配自己。 查询时自己还没放进去，
		// 所以找到的一定是另一个位置，不用再写 j != i。
		//
		// 在当前元素之前的部分找另一个数，而不是之后的部分找另一个数。
		// 当前元素之前的部分保存在 map 中
		y := target - x
		if j, found := numMap[y]; found {
			return []int{i, j}
		}
		numMap[x] = i
	}

	return nil
}

// 暴力解法：用两层循环遍历所有的数对。
// 外层循环迭代第一个元素的候选列表，内层循环遍历第二个元素的候选列表
// 运行会超时
func twoSumVariant2_v4(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	for i, v1 := range nums {
		for j := i + 1; j < n; j++ {
			if v1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
