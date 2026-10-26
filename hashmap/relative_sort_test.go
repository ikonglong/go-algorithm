package hashmap

import "slices"

// https://leetcode.cn/problems/relative-sort-array/
// 算法思路：
// 1. 既然要将 a1 中的元素排序，使得 a1 中在 a2 中的元素按 a2 的相对顺序排列，那么首先得算出 a1 中哪些元素在 a2
// 2. 考虑为 a2 创建 a2_map，key 为元素值，value 待定
// 3. 考虑到 a1 中的元素可能重复，例如 a1、a2 都包含 2，但 a1 中可能有多个 2，
//    因此将 a2_map 的 value 定义为 a2 元素在 a1 中的重复次数。这样 a2_map 就可以发挥更大价值。
//    通常，一个 map 的 key 和 value 来自同一个源对象；但如果打破这个思考惯性，让 key 和 value 来自不同的源对象，
// 	  很可能让 map 发挥更大价值
// 4. 遍历 a1，将 a1 中存在于 a2 的元素添加至结果列表
// 5. ...

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// key 为 a2 元素，value 为该元素在 a1 的重复次数
	a2Map := make(map[int]int, len(arr2))
	for _, v := range arr2 {
		a2Map[v] = 0
	}
	// 计算元素在 a1 的重复次数
	for _, v := range arr1 {
		if count, found := a2Map[v]; found {
			a2Map[v] = count + 1
		}
	}
	// 将 a2 中存在于 a1 的元素按照 a2 的顺序添加至结果数组
	result := make([]int, len(arr1), len(arr1))
	i := 0
	for _, item := range arr2 {
		// for _ = range a2Map[item] {
		// 	result[i] = item
		// 	i = i + 1
		// 	// bug: j=1，i=1，设置result[1]后期望i=2，但i=1+1+1=3
		// 	//i = i + j + 1
		// }

		// 这样写也 OK
		count := a2Map[item]
		for j := range count {
			result[i+j] = item
		}
		i += count
	}
	// 将 a1 中剩余元素按照升序添加至结果数组
	// 找出 a1 中不在 a2 的元素列表
	absent := make([]int, 0)
	for _, v := range arr1 {
		if _, found := a2Map[v]; !found {
			absent = append(absent, v)
		}
	}
	slices.Sort(absent)
	for _, v := range absent {
		result[i] = v
		i++
	}
	return result
}
