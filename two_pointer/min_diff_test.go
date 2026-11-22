package twopointer

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/smallest-difference-lcci/
// 计算最小差绝对值
//
// 算法思路，类似合并两个有序数组：
// 1. 排序两个数组
// 2. 创建两个指针 pa、pb，分别指向两个数组的头元素
// 3. 使用 pa、pb 遍历两个数组，直到任何一个数组遍历完
// 		a. 计算 pa、pb 所指元素的差绝对值，若小于全局结果，则更新全局结果。这就是循环体的核心逻辑
// 		b. 接下来为下个迭代做准备
//      比较两个元素，根据「找出最小差」这个目标，决定移动哪个指针
//      a[pa] > b[pb], pb++
//      a[pa] < b[pb], pa++
//      a[pa] == a[pb]，差值为 0，肯定是最小差绝对值了，退出循环，直接返回

func smallestDifference_(a []int, b []int) int {
	slices.Sort(a)
	slices.Sort(b)
	pa, pb := 0, 0
	min := diffAbs(a[pa], b[pb])

	for pa < len(a) && pb < len(b) {
		diff := diffAbs(a[pa], b[pb])
		if diff < min {
			min = diff
		}

		// 为下个迭代做准备
		if a[pa] > b[pb] {
			pb++
		} else if a[pa] < b[pb] {
			pa++
		} else {
			return 0
		}
	}

    return min
}

func diffAbs(n1, n2 int) int {
	return int(math.Abs(float64(n1-n2)))
}
