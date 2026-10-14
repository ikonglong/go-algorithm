package binary_search

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/sparse-array-search-lcci/
// 稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
// 例如：words=["at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""], s="ta"
// 注意，空字符串的存在使得二分搜索失效了
//
// 对["b", "", " ", "a"]排序，结果为：["", " ", "a", "b"]

// 算法思路有两种：
// 在二分查找循环体中：
// 1. 若 words[mid] 为 ""，就向右或向左移动 mid，直到找到一个非空的值
// 2. 若 words[mid] 为 ""，就移动 low 或 high，但要判断当前 low 或
//    high 指向的元素是否是目标值，若是则直接返回。这个思路通过移动 low 或 high 改变 mid
// 这两种思路在最坏情况下的时间复杂度都是 O(n)

// 若 words[mid] 为空字符串，则移动 low
func findString(words []string, s string) int {
	lo, hi := 0, len(words)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if words[mid] == s {
			return mid
		} else if words[mid] == "" {
			if words[lo] == s {
				return lo
			}
			lo++
		} else if words[mid] < s {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// 虽然能工作，但这个实现很是脱裤子放屁，因为构建了单词到下标的映射，就不需要二分查找了
func findString3(words []string, s string) int {
	// 既然都已经构建单词到下标的映射了，就不需要二分查找了
	wordIndexes := make(map[string]int)
	for i, v := range words {
		if v != "" {
			wordIndexes[v] = i
		}
	}
	slices.Sort(words)
	low, high := 0, len(words)-1
	for i, v := range words {
		if v != "" {
			low = i
			break
		}
	}
	for low <= high {
		mid := (low + high) / 2
		if words[mid] == s {
			// return mid
			// 既然都已经构建单词到下标的映射了，就不需要二分查找了
			return wordIndexes[s]
		} else if words[mid] < s {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func TestFindWord(t *testing.T) {
	r := findString3(
		[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""},
		"ball")
	assert.Equal(t, 4, r)
}
