package sort_

import "sort"

// https://leetcode.cn/problems/merge-intervals/
// 该题跟 attend_meeting_test.go 类似

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for _, item := range intervals {
		n := len(result)
		if n == 0 {
			result = append(result, item)
			continue
		}
		prev := result[n-1]
		if item[0] <= prev[1] { // 有重叠
			// []int{min(prev[0], item[0]), max(prev[1], item[1])}
			result[n-1] = []int{prev[0], max(prev[1], item[1])}
		} else {
			result = append(result, item)
		}
	}
	return result
}
