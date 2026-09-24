package twopointer

import "math"

// https://leetcode.cn/problems/find-closest-lcci/description/
// 分析题目：
// 找出两个单词 A、B 之间的最短距离，不要死板地认为一定 A 在前、B 在后。
// A、B 之间的距离等价于 B、A 之间的距离

// 算法思路：
// 1. 创建两个列表 w1Pos、w2Pos，分别保存 word1、word2 在原字符串中的位置
// 2. 采用类似合并、归并两个有序数组的方法。创建两个指针 p1、p2，分别指向两个列表的头元素。重复下述步骤：
// 		a. 如果 w1Pos[p1] > w2Pos[p2]，先计算最小距离，再将较小的头元素之中的右移；否则，亦如此。
//
// 头元素较小，意味着该单词在原始字符串中的位置靠左。
// 举个例子：[student, a, a, a, student]
// w1ps: [0, 4]
// w2ps: [1, 2, 3]
// 根据题目，虽然 1-4 是合法的配对，但求最小距离时会被自然淘汰

func findClosest(words []string, word1 string, word2 string) int {
	w1Pos := make([]int, 0)
	w2Pos := make([]int, 0)
	for i, v := range words {
		if v == word1 {
			w1Pos = append(w1Pos, i)
		} else if v == word2 {
			w2Pos = append(w2Pos, i)
		}
	}

	minDistance := math.MaxInt
	for p1, p2 := 0, 0; p1 < len(w1Pos) && p2 < len(w2Pos); {
		if w1Pos[p1] > w2Pos[p2] {
			minDistance = min(minDistance, w1Pos[p1]-w2Pos[p2])
			p2++
		} else if w2Pos[p2] > w1Pos[p1] {
			minDistance = min(minDistance, w2Pos[p2]-w1Pos[p1])
			p1++
		}
	}
	return minDistance
}

// 算法思路：把这个问题转化为：在 words 数组分别寻找两个单词 a、b，
// 找到 a、b 位置差值绝对值的最小值。
// 逻辑上可以认为创建两个指针，分别用来在数组上查找两个目标单词

func findClosest2(words []string, word1 string, word2 string) int {
	minDistance := math.MaxInt
	n := len(words)
	prev, curr := -1, 0
	prevWord := ""
	for ; curr < n; curr++ {
		w := words[curr]
		if w == word1 || w == word2 {
			if prevWord != "" && w != prevWord {
				minDistance = min(curr-prev, minDistance)
			}
			prev = curr
			prevWord = w
		}
		// 埋有 bug：连续遇到同一目标词时，也必须更新 prev、prevWord，但下面实现思路未更新。
		// 例如：...,"a", "a", "student",...
		// found := (prevWord == "" && (words[curr] == word1 || words[curr] == word2)) ||
		// 	(prevWord == word1 && words[curr] == word2) ||
		// 	(prevWord == word2 && words[curr] == word1)
		// if found {
		// 	if prev != -1 {
		// 	}
		// 	prev = curr
		// 	prevWord = words[curr]
		// }
	}
	return minDistance
}

func findClosest3_bug(words []string, word1 string, word2 string) int {
	minDistance := math.MaxInt
	n := len(words)
	prev, curr := -1, 0
	for ; curr < n; curr++ {
		// 埋有 bug：两个连续的 word1 或 word2
		// 被错误地认为是 word 1 和 word 2 的搭配
		if words[curr] != word1 || words[curr] != word2 {
			if prev != -1 {
				minDistance = min(curr-prev, minDistance)
			}
			prev = curr
		}
	}
	return minDistance
}
