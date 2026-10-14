package sort_

import "sort"

// https://leetcode.cn/problems/meeting-rooms/description/
// 给定一个会议时间安排的数组 intervals ，其中 intervals[i] = [starti, endi]。
// 一个人可以参加所有会议，只要没有两个会议的时间段重叠。在时间 t 结束的会议和在时间 t 开始的会议不重叠。
// 如果一个人能够参加这里面的全部会议，返回 true，否则返回 false。

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 按会议的开始时间对所有会议排序后，若后一个会议的开始时间
	// 小于前一个会议的结束时间，则无法参加全部会议
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
