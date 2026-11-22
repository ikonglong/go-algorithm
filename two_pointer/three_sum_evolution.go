package twopointer

import "slices"

// https://leetcode.cn/problems/3sum/

// 在下一个版本的基础上，实现了全局去除
func threeSum2_(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	for i, a := range nums {
		// 对第一个数去重
		if i > 0 && a == nums[i-1] {
			continue
		}

		pairs := pairNums(nums[i+1:], -a) // ok 的
		// pairs := pairNums2(nums[i+1:], -a)
		for _, p := range pairs {
			result = append(result, append([]int{a}, p...))
		}
	}
	return result
}

// 在下一个版本上的改进：求两个数的和时，不再用暴力解法，将暴力解法替换为双指针
// 最终结果中仍有重复答案
func threeSum3_(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	// n := len(nums)
	for i, a := range nums {
		// 虽然 pairNums 调用内部会去重，但它只能去掉一次调用中发现的重复，
		// 多次 pairNums 调用返回的答案列表间仍可能有重复。如何全局去重？
		pairs := pairNums(nums[i+1:], -a)
		for _, p := range pairs {
			result = append(result, append([]int{a}, p...))
		}

		// bug: 确定第一个数后，就可直接在子数组中查找和为 -a 的所有数对了。
		// 不需要，也不应该还像暴力解法那样，在每一个子数组中查找数对。
		//
		// 之所以会犯这么脑残的低级错误，就是因为实现当前思路时直接拷贝了暴力解法的代码作为代码骨架，
		// 在此基础上删减修改，导致思考路径被暴力解法的代码骨架框死了，带跑偏了。
		// 暴力解法的思考框架和代码骨架，并不适合作为当前算法思路的思考框架和代码骨架。
		//
		// 避免掉进这种坑里的最佳实践是：不要偷懒，想清楚算法思路后，从零开始写代码构架
		//
		// for j := i + 1; j < n; j++ {
		// 	pairs := pairNums(nums[j:], -a)
		// 	for _, p := range pairs {
		// 		result = append(result, append([]int{a}, p...))
		// 	}
		// }
	}
	return result
}

func pairNums(nums []int, target int) [][]int {
	// 外层已经全局排序了，这里就不用再将 nums 排序了
	result := make([][]int, 0)
	lo, hi := 0, len(nums)-1
	lastAns := []int{}
	for lo < hi {
		s := nums[lo] + nums[hi]
		if s == target {
			// 用 lastAns 去重没错，但它只能去掉一次 pairNums 调用中发现的重复。
			// 每次调用都会重新初始化 lastAns，如果 lastAns 是全局持续维护的，是不是就能做到全局去重？
			// 有更好的去重办法，不在这个思路上继续想了
			if len(lastAns) == 0 || nums[lo] != lastAns[0] || nums[hi] != lastAns[1] {
				ans := []int{nums[lo], nums[hi]}
				result = append(result, ans)
				lastAns = ans
			}
			lo++
			hi--
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}
	return result
}

// 在上一个 pairNums 实现版本上改进去重。
// 不再比较后两个数组成的数对，而是通过先对第一个数去重，再对第二个数去重，来实现全局去重。
// 对第一个数的去重在 pair_nums 的调用函数中做
func pairNums2(nums []int, target int) [][]int {
	// 外层已经全局排序了，这里就不用再将 nums 排序了
	result := make([][]int, 0)
	lo, hi := 0, len(nums)-1
	for lo < hi {
		// 对答案的第二个数去重
		if lo > 0 && nums[lo] == nums[lo-1] {
			lo++
			continue
		}

		s := nums[lo] + nums[hi]
		if s == target {
			result = append(result, []int{nums[lo], nums[hi]})
			lo++
			hi--
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}
	return result
}

// 暴力解法，超出时间限制
func threeSum4_(nums []int) [][]int {
	result := make([][]int, 0)
	aMap := make(map[[3]int]struct{})
	n := len(nums)
	for i, v := range nums {
		a := v
		for j := i + 1; j < n; j++ {
			b := nums[j]
			c := 0 - a - b
			for k := j + 1; k < n; k++ {
				if nums[k] == c {
					answer := []int{a, b, c}
					slices.Sort(answer)
					sortedAns := [3]int(answer)
					if _, found := aMap[sortedAns]; !found {
						result = append(result, answer)
						aMap[sortedAns] = struct{}{}
					}
				}
			}
		}
	}
	return result
}
