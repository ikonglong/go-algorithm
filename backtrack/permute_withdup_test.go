package backtrack

import (
	"slices"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	permuts := make(map[string]struct{})

	var backtrack func(k int, path []int, idxPath []int)
	backtrack = func(k int, path []int, idxPath []int) {
		if k == len(nums) {
			key := setToStr(path)
			if _, found := permuts[key]; !found {
				permuts[key] = struct{}{}
				result = append(result, append([]int(nil), path...))
			}
			return
		}

		for i, v := range nums {
			if slices.Contains(idxPath, i) {
				continue
			}

			path = append(path, v)
			idxPath = append(idxPath, i)
			backtrack(k+1, path, idxPath)
			path = path[:len(path)-1]
			idxPath = idxPath[:len(idxPath)-1]
		}
	}

	backtrack(0, []int(nil), []int(nil))
	return result
}

func setToStr(set []int) string {
	var b strings.Builder
	for _, v := range set {
		b.WriteString(strconv.Itoa(v))
		b.WriteRune(',')
	}
	return b.String()
}
