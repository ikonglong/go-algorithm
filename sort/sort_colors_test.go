package sort_

import "slices"

// https://leetcode.cn/problems/sort-colors/description/
// 整数 0、 1 和 2 分别表示红色、白色和蓝色。【原地】对它们进行排序，
// 使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 这里采用排序思路实现，其他思路请看 two_pointer/sort_colors_test.go
// 但 leetcode 禁止使用 stdlib sort 函数：
// Use of slices.Sort() is a prohibited built-in sorting operation.
func sortColors(nums []int) {
	slices.Sort(nums)
}
