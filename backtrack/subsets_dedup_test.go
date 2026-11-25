package backtrack

import (
	"slices"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/subsets-ii/
//
// 需要关注是否有序
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	// 因为输入集合中存在重复数字，导致产生的组合会有重复，利用这个 set 去重
	subsets := make(map[string]struct{})

	// 输入集合中有多少个数字，就有多少个决策阶段
	// 每个阶段处理一个数字，有两种选择：添加、不添加
	// 路径，已完成阶段选择的数字组成的组合
	var backtrack func(k int, comb []int)
	backtrack = func(k int, comb []int) {
		if k == len(nums) {
			key := combToStr(comb)
			if _, found := subsets[key]; !found {
				subsets[key] = struct{}{}
				result = append(result, append([]int(nil), comb...))
			}
			return
		}

		// 不添加
		backtrack(k+1, comb)

		// 添加
		comb = append(comb, nums[k])
		backtrack(k+1, comb)
		comb = comb[:len(comb)-1]
	}

	backtrack(0, []int(nil))
	return result
}

func combToStr(comb []int) string {
	var b strings.Builder
	for _, v := range comb {
		// 上一行代码埋有 bug，不加元素后不加逗号时，有可能两个完全不同的组合生成了相同的字符串
		// fmt.Fprintf(&b, "%d", v)

		// fmt.Fprintf(&b, "%d,", v)
		b.WriteString(strconv.Itoa(v))
		b.WriteRune(',')
	}
	return b.String()
}
