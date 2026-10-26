package hashmap

// https://leetcode.cn/problems/intersection-of-two-arrays/

// 算法思路：
// 借助哈希表实现就比较简单，创建一个 map，遍历 nums1，将其中的元素放入 map。
// 再遍历 nums2，若元素在 map 中，将其添加至结果列表
//
// 其他实现思路：binary_search/intersection_test.go

func intersection(nums1 []int, nums2 []int) []int {
	aMap := make(map[int]struct{})
	for _, item := range nums1 {
		aMap[item] = struct{}{}
	}
	// estimLen := len(nums1)
	// if len(nums1) > len(nums2) {
	// 	estimLen = len(nums2)
	// }
	estimLen := min(len(nums2), len(nums1))
	result := make([]int, 0, estimLen)
	// 为了去重重复的交集元素
	intersec := make(map[int]struct{}, estimLen)
	for _, item := range nums2 {
		if _, found := aMap[item]; found {
			// result = append(result, item)
			intersec[item] = struct{}{}
		}
	}
	for k, _ := range intersec {
		result = append(result, k)
	}
	return result
}
