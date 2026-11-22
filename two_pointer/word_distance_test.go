package twopointer

import "math"

// https://leetcode.cn/problems/find-closest-lcci/
// 分析题目：
// 找出两个单词 w1、w2 之间的最短距离，不要死板地认为一定是 w1 在前，w2 在后。
// w1、w2 之间的距离等价于 w2、w1 之间的距离
//
// 举个例子：[student, a, a, a, student]
// w1Pos: [0, 4]
// w2Pos: [1, 2, 3]
// 根据题目，虽然 1-4 是合法的配对，但求的是最小距离，因此被自然淘汰

// 算法思路：
// 1. 创建两个列表 w1Pos、w2Pos，分别保存 w1、w2 在原字符串中的位置
// 2. 有了这两个列表，原问题 就转化为 求两个整型数组元素之间的最小差绝对值了，
//    算法就不再详述了，看 two_pointer/min_diff_test.go

func findClosest(words []string, word1 string, word2 string) int {
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
		// 上面是循环体的核心逻辑

		// 接下来为下个迭代做准备
		// 处理逻辑类似归并两个有序数组：当前的两个元素中哪个小，其所属数组的指针就右移
		if w1Pos[i1] > w2Pos[i2] {
			i2++
		} else { // 因为不可能相等，所以等价于 w1Pos[i1] < w2Pos[i2]
			i1++
		}
	}
	return min
}

// 这个版本采用的算法思路跟上一个版本完全相同。实现上，除了主循环中「为下个迭代做准备」
// 的逻辑换了一种等价的写法外，没有其他不同。这么做是为了呈现出这部分逻辑跟归并两个
// 有序数组时采用的思路类似

func findClosest2(words []string, word1 string, word2 string) int {
	// 构造两个数组，分别保存两个单词在文件中的出现位置
	w1Pos := make([]int, 0)
	w2Pos := make([]int, 0)
	for i, v := range words {
		if v == word1 {
			w1Pos = append(w1Pos, i)
		} else if v == word2 {
			w2Pos = append(w2Pos, i)
		}
	}

	min := math.MaxInt
	for i1, i2 := 0, 0; i1 < len(w1Pos) && i2 < len(w2Pos); {
		dist := int(math.Abs(float64(w1Pos[i1]-w2Pos[i2])))
		if dist < min {
			min = dist
		}
		// 上面是循环体的核心逻辑

		// 接下来为下个迭代做准备
		// 处理逻辑类似归并两个有序数组：当前的两个元素中哪个小，其所属数组的指针就右移
		if w1Pos[i1] < w2Pos[i2] {
			i1++
		} else if w2Pos[i2] < w1Pos[i1] {
			i2++
		}
	}
	return min
}
