package twopointer

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/smallest-difference-lcci/
// 计算最小差绝对值
//
// 算法思路：
// 1. 排序两个数组
// 2. 创建两个指针 pa、pb，分别指向两个数组的头元素
// 3. 比较两个指针，以根据“找出最小差”的目标，决定移动哪个指针
//      pa > pb, pb++
//      pa < pb, pa++
//      pa == pb, 差值绝对值为 0，肯定是最小差值绝对值，退出循环，直接返回
//      任何一个数组遍历完就结束循环

func smallestDifference(a []int, b []int) int {
	result := math.MaxInt
	slices.Sort(a)
	slices.Sort(b)

	pa, pb := 0, 0
	for pa < len(a) && pb < len(b) {
		diffAbs := int(math.Abs(float64(a[pa] - b[pb])))
		if a[pa] > b[pb] {
			pb++
		} else if a[pa] < b[pb] {
			pa++
		} else {
			return 0
		}
		if diffAbs < result {
			result = diffAbs
		}
	}

	return result
}
