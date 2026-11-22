package twopointer

import "math"

// https://leetcode.cn/problems/find-closest-lcci/

func findClosest_(words []string, word1 string, word2 string) int {
	// 构造两个数组，分别保存两个单词在文件中的出现位置
	w1Pos := make([]int, 0)
	w2Pos := make([]int, 0)
	for i, w := range words {
		if w == word1 {
			w1Pos = append(w1Pos, i)
		} else if w == word2 {
			w2Pos = append(w2Pos, i)
		}
	}

	min := math.MaxInt
	for i1, i2 := 0, 0; i1 < len(w1Pos)&& i2 < len(w2Pos); {
		dist := int(math.Abs(float64(w1Pos[i1]-w2Pos[i2])))
		if dist < min {
			min = dist
		}

		if w1Pos[i1] > w2Pos[i2] {
			i2++
		} else { // 因为不可能相等，所以等价于 w1Pos[i1] < w2Pos[i2]
			i1++
		}
	}
	return min
}
