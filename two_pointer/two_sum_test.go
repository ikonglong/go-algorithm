package twopointer

import "slices"

// https://leetcode.cn/problems/pairs-with-sum-lcci/
// 注意，一个元素只能使用一次。返回满足要求的元素。

// 算法：
//  1. 先排序
//  2. 创建两个指针 low、high，分别指向第一个、最后一个元素
//  3. 使用这两个指针遍历数组，直到 low 大于等于 high
//     a. 计算 low、high 所指元素之和，记为 s
//     b. 若 s 等于 target，将这对数添加到结果中，再右移 low，左移 high
//     c. 若大于，则向左移动 high
//     c. 若小于，向右移动 low
//
// 容易犯的错误：相等时只左移 hi
// 对于输入[5,5,6,6]，若相等时只左移 hi，就会错过另一对合法结果
// 因为一个元素只能用一次，所以相等时必须 low 右移，high 左移

// 基于下一个版本改进出来的更简洁易读的版本
func pairSums_(nums []int, target int) [][]int {
	result := make([][]int, 0)
	slices.Sort(nums)
	lo, hi := 0, len(nums)-1
	for lo < hi {
		s := nums[lo] + nums[hi]
		if s == target {
			result = append(result, []int{nums[lo], nums[hi]})
			// 容易犯的错误：相等时只左移 hi
			// 对于输入[5,5,6,6]，若相等时只左移 hi，就会错过另一对合法结果
			// 因为一个元素只能用一次，所以相等时必须 low 右移，high 左移
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

// 基于下一个有 bug 的版本改进而来，能工作但还不够简洁易读
func pairSums2_(nums []int, target int) [][]int {
	result := make([][]int, 0)
	slices.Sort(nums)
	lo, hi := 0, len(nums)-1
	// 最后一个结果中元素的索引
	lastLo, lastHi := -1, -1
	for lo < hi {
		s := nums[lo] + nums[hi]
		if s == target {
			// 去重。注意，允许不同元素的值相同，但不允许使用同一个元素多次
			if lastLo == lo || lastHi == hi {
				// do nothing
			} else {
				result = append(result, []int{nums[lo], nums[hi]})
				lastLo = lo
				lastHi = hi
			}

			// 前面的 if 块简洁一点的写法
			// if lastLo != lo && lastHi != hi {
			// 	result = append(result, []int{nums[lo], nums[hi]})
			// 	lastLo = lo
			// 	lastHi = hi
			// }

			// 容易犯的错误：相等时只左移 hi
			// 对于输入[5,5,6,6]，若相等时只左移 hi，就会错过另一对合法结果
			// 因为一个元素只能用一次，所以相等时必须 low 右移，high 左移
			//
			// 既然 low、high 都必须移动，自然就不会产生重复的结果了，那么去重逻辑是不是可以删除了
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

// 有多个 bug，但这些 bug 的根因都分析清楚了，可以基于它改进出可工作的版本
func pairSums_bad(nums []int, target int) [][]int {
	result := make([][]int, 0)
	slices.Sort(nums)
	lo, hi := 0, len(nums)-1
	last:=[]int{}
	for lo < hi {
		s := nums[lo] + nums[hi]
		if s == target {
			pair := []int{nums[lo], nums[hi]}
			// 去重
			// bug: 一个元素只能用一次，但不同位置的元素可以相同，进而值相同但位置不同的元素并不是重复结果呀！
			// if len(result) == 0 || (len(last) == 2 && last[0] != pair[0] && last[1] != pair[1]) {
			// bug: 跟最后一个结果不重复时要添加，但结果列表为空时也要添加呀！
			// if len(last) == 2 && last[0] != pair[0] && last[1] != pair[1]{
			// 两个结果，只要其中一个数相同，另一个必然相同，因此只需检测一个是否相同
			// bug: 跟最后一个结果不重复时要添加，但结果列表为空时也要添加呀！
			if len(last) == 2 && last[0] != pair[0] {
				result = append(result, pair)
				last = pair
				// bug: 不管结果是否重复，hi 都要左移
				hi--
			}
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}
	return result
}
