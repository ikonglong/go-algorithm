package twopointer

import (
	"fmt"
	"slices"
)

// https://leetcode.cn/problems/two-sum/
// 注意，一个元素只能使用一次，每个数只有一个答案。返回满足要求的元素的下标

// 推荐的实现版本：twoSumVariant2、twoSumVariant23

// 在给定数组中寻找和为 target 的数对。有两种方法分解问题：
// 1. 既然是两数 a、b 之和，那么另一个数就是 target-a
// 2. 先用某种方式将数组划分成一个个数对；再计算和，跟 target 比较

// 双指针思路，详细描述请看 two_pointer/two_sum_test.go
// 这道题特别的地方：返回答案的下标。但是算法会先对原数组排序，下标就变了。
// 因此需要思考清楚如何保存原始下标，考虑使用 map 保存。但如果答案中的两个
// 元素值相同，后一个的下标就覆盖了前一个。如何解决？
// 用一个长度相等的布尔数组记录哪个下标对应的元素已经使用了，但找到答案时拿
// 到的是排序后的下标，还得拿对应元素到原始数组中找到原下标，检查是否使用了，
// 若已使用，继续找下一个相等的元素的下标

// 不再创建一个标记数组来记录下标对应的原始元素是否已使用
// 只把已使用的元素的下标保存到一个列表中
func twoSumVariant2(nums []int, target int) []int {
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
				origIdx(sorted[lo], nums, &used),
				origIdx(sorted[hi], nums, &used),
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
func origIdx(v int, nums []int, used *[]int) int {
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

// 对于此题来说，只是为了记录一个下标是否使用了，就创建了一个标记数组，没必要，可优化
func twoSumVariant22(nums []int, target int) []int {
	// 用 map 保存原始下标行不通，因为答案中的两个元素可能值相同
	// idxMap := make(map[int]int)
	// for i, e := range nums {
	// 	idxMap[e] = i
	// }

	// 针对答案中两个元素相同的 case，例如：输入 [3,3]
	// 用一个长度相等的布尔数组记录哪个下标在答案中已经使用了
	usedFlags := make([]bool, len(nums))
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
				origIdx2(sorted[lo], nums, usedFlags),
				origIdx2(sorted[hi], nums, usedFlags),
			}
		} else if s > target {
			hi--
		} else {
			lo++
		}
	}
	return nil
}

// usedFlags slice 是传值调用，但在这里是可以的，因为它在这里的
// 本质是数组，创建时所有元素都已初始化，只需要修改元素值，没有扩缩操作
func origIdx2(v int, nums []int, usedFlags []bool) int {
	for i, e := range nums {
		if v == e && !usedFlags[i] {
			usedFlags[i] = true
			return i
		}
	}
	return -1 // or panic error("not found %d", v)
}

// 先确定一个数，再求出另一个数，然后在哈希表中查找另一个数的下标
//
// 算法思路：
//  1. 创建一个 map，键是元素值，值是下标，用来保存已遍历过的元素的下标和值
//  2. 假设当前遍历到 e。
//     在 map 中找 target-e。若找到，直接返回答案；若未找到，将 e 添加至 map，继续遍历
//
// 这个思路很巧妙，如果遍历到 e 时先把 e 放入 map，而 target-e 又刚好等于 e 的话，从
// map 找到的下标是 e 的下标，而不是 target-e 的。实际上遍历到 e 时，e 的下标已经在手，
// 只需要找到 target-e 的下标就够了，完全没必要把 e 放入 map
func twoSumVariant23(nums []int, target int) []int {
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

// 先确定一个数，再求出另一个数，然后在哈希表中查找另一个数的下标
func twoSumVariant24(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	aMap := make(map[int]int, len(nums))
	for i, v := range nums {
		aMap[v] = i
	}

	// 若有元素值相同，构造 map 时后面的覆盖前面的，也不影响这个算法正常工作。
	// 因为题目说返回任何一个答案即可，例如输入 [3,3,4]，target=7，返回 [1,2] 是 OK 的

	for i, x := range nums {
		// 想通过删除自己来避免：当两个数相同时，原本要找另一个数的下标，结果找到的是自己的下标
		// delete(aMap, x)

		y := target - x
		// 通过 i != j 避免自己跟自己配对
		// 粗想时我认为这么做并不能实现想要的效果。例如：输入 [3, 3]，target=6，能否得到正确结果？
		// 我认为结果是 [1,1]，不正确。但实际验证能得到正确结果 [0,1]。因为虽然构建 map 时后面的覆
		// 盖前面的，map:{3:1}，但查找时仍然遍历整个数组，这就使得位于前面的被覆盖的元素依然有机会跟
		// 后面的留存在 map 中的元素配对
		if j, found := aMap[y]; found && i != j {
			return []int{i, j}
		}
	}
	return nil
}

// 暴力解法：用两层循环遍历所有的数对。
// 外层循环迭代第一个元素的候选列表，内层循环遍历第二个元素的候选列表
func twoSumVariant25(nums []int, target int) []int {
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
