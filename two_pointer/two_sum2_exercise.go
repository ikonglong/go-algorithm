package twopointer

import (
	"fmt"
	"slices"
)

// https://leetcode.cn/problems/two-sum/
// 注意，一个元素只能使用一次，每个数只有一个答案。返回满足要求的元素的下标

// 双指针思路，详细描述请看 two_pointer/two_sum_test.go
// 这道题特别的地方：返回答案的下标。但是算法会先对原数组排序，下标就变了。
// 因此需要思考清楚如何保存原始下标，考虑使用 map 保存。但如果答案中的两个
// 元素值相同，后一个的下标就覆盖了前一个。如何解决？
// 用一个长度相等的布尔数组记录哪个下标对应的元素已经使用了，但找到答案时拿
// 到的是排序后的下标，还得拿对应元素到原始数组中找到原下标，检查是否使用了，
// 若已使用，继续找下一个相等的元素的下标

// 不再创建一个标记数组来记录下标对应的原始元素是否已使用
// 只把已使用的元素的下标保存到一个列表中
func twoSumVariant2_(nums []int, target int) []int {
	// 用 map 保存原始下标行不通，因为答案中的两个元素可能值相同
	// idxMap := make(map[int]int)
	// for i, e := range nums {
	// 	idxMap[e] = i
	// }

	// 针对答案中两个元素相同的 case，例如：输入 [3,3]
	// 保存在答案中已使用的原始下标。也可以 used := [1]int{-1}
	used := make([]int, 0)
	// 排序很重要，不要漏了，漏了就不对了
	// slices.Sort(nums)
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	slices.Sort(sorted)

	lo, hi := 0, len(sorted)-1
	for lo < hi {
		s := sorted[lo] + sorted[hi]
		if s == target {
			// return []int{lo, hi}
			// return []int{idxMap[nums[lo]], idxMap[nums[hi]]}
			return []int{
				origIdx_(sorted[lo], nums, &used),
				origIdx_(sorted[hi], nums, &used),
			}
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}
	return nil
}

// 这里 used []int 是传值调用，而第一次调用会扩 slice，传值调用
// 的效果是调用内部扩了 used，调用返回后调用方 used 依然没变。
// 因此这里必须用指针
func origIdx_(v int, nums []int, used *[]int) int {
	usedV := *used
	for i, e := range nums {
		if v == e && (len(usedV) == 0 || usedV[0] != i) {
			// bug: 并未修改指针所指向的变量
			// usedV = append(usedV, i)
			*used = append(usedV, i)
			return i
		}
	}
	panic(fmt.Errorf("not found %d", v))
}

// 算法思路：
//  1. 创建一个 map，键是元素值，值是下标，用来保存已遍历过的元素的下标和值
//  2. 假设当前遍历到 e。
//     在 map 中找 target-e。若找到，直接返回答案；若未找到，将 e 添加至 map，继续遍历
//
// 这个思路很巧妙，如果遍历到 e 时先把 e 放入 map，而 target-e 又刚好等于 e 的话，从
// map 找到的下标是 e 的下标，而不是 target-e 的。实际上遍历到 e 时，e 的下标已经在手，
// 只需要找到 target-e 的下标就够了，完全没必要把 e 放入 map
func twoSumVariant23_(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	aMap := make(map[int]int)
	for i, e := range nums {
		if idx, found := aMap[target-e]; found {
			return []int{idx, i}
		}
		aMap[e] = i
	}
	return nil
}
