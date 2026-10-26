package hashmap

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// 找到重复元素，返回任一副本

// 算法思路：
// 1. 创建一个 map，key 为 元素值，value 为空。
// 2. 遍历数组：若当前元素不存在于 map，则将其放入 map；若存在，则返回当前元素
//
// 可否用双指针解？

func findRepeatDocument(documents []int) int {
	aMap := make(map[int]struct{})
	for _, v := range documents {
		if _, found := aMap[v]; found {
			return v
		} else {
			aMap[v] = struct{}{}
		}
	}
	return -1
}
