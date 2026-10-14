package sort_

import "sort"

// https://leetcode.cn/problems/can-make-arithmetic-progression-from-sequence/description/
// 判断是否是等差数列
// 两个思路：
// 判断相邻两项的差是否相等
// 判断 2arr[i] = arr[i-1] + arr[i+1]

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) < 2 {
		return false
	}
	diff := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if diff != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
