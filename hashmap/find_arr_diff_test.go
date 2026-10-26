package hashmap

import "slices"

// https://leetcode.cn/problems/find-the-difference-of-two-arrays/
// 返回 nums1 中不在 nums2 的元素列表，以及 nums2 中不在 nums1 的元素列表
// 这道题的另一种等价描述：从 nums1 移除跟 nums2 的重复元素，仅保留不重复元素。
// 不过这道题要求两个结果：
// 1. 以 nums1 为基准集合或源集合
// 2. 以 nums2 为基准集合或源集合

// 算法思路：
// 1. 这道题求两个结果，分别以 nums1、nums2 为源集合或基准集合。因此不能在源数组
//    上直接执行删除操作。
// 2. 为两个求解分别创建一个结果数组
// 3. 先求解以 nums1 为原集合，求在 nums1 中但不在 nums2 的元素列表。
// 		a. 创建一个 map，key 为元素值，value 为空，遍历数组 2 初始化 map
// 		b. 遍历数组 1，若当前元素不在 map，添加到结果数组
// 4. 用相同的思路做第二个求解

func findDifference(nums1 []int, nums2 []int) [][]int {
	return [][]int{
		doDiff(nums1, nums2, 1),
		doDiff(nums1, nums2, 2),
	}
}

// source: 1 -> nums1, 2 -> nums2
func doDiff(nums1 []int, nums2 []int, source int) []int {
	var srcSet, exclSet []int
	if source == 1 {
		srcSet = nums1
		exclSet = nums2
	} else {
		srcSet = nums2
		exclSet = nums1
	}

	exclMap := make(map[int]struct{})
	for _, v := range exclSet {
		exclMap[v] = struct{}{}
	}

	result := make([]int, 0)
	slices.Sort(srcSet)
	for i, v := range srcSet {
		// 若当前元素跟前一个元素重复，则不管其是否已加入结果列表都跳过，
		// 这样就可去除结果列表中的重复元素
		if i > 0 && srcSet[i-1] == v {
			continue
		}
		if _, found := exclMap[v]; !found {
			result = append(result, v)
		}
	}

	return result
}
