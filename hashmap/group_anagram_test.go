package hashmap

import "slices"

// https://leetcode.cn/problems/group-anagrams/description/
// 如何判断两个字符串是否是异位词，已经挺熟悉了。这里主要思考如何分组。
// 算法思路：
// 1. 创建一个 Map，用来对彼此是异位词的字符串进行分组
// 2. 遍历字符串数组：
// 若当前元素在 map 中存在，则将其添加至对应列表；若不存在，检查排序后的值是否存在。
// 若存在，将原始值添加至对应列表；否则，将 (key:排序后的值, value:包含原始值的列表)
// 放入 map

func groupAnagrams(strs []string) [][]string {
	aMap := make(map[string][]string)
	for _, s := range strs {
		if l, found := aMap[s]; found {
			aMap[s] = append(l, s)
			continue
		}
		chars := []byte(s)
		slices.Sort(chars)
		sortedS := string(chars)
		if l, found := aMap[sortedS]; found {
			aMap[sortedS] = append(l, s)
		} else {
			aMap[sortedS] = append([]string(nil), s)
		}
	}
	result := make([][]string, len(aMap))
	i := 0
	for _, l := range aMap {
		result[i] = l
		i++
	}
	return result
}
