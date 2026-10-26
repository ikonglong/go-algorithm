package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/single-number/
// 数组中只有一个数出现一次，其他数均出现两次，找到那个只出现一次的数

func singleNumber(nums []int) int {
	// key 元素值，value 重复次数
	aMap := make(map[int]int)
	for _, v := range nums {
		aMap[v] = aMap[v] + 1
	}
	// 多么愚蠢的 bug：把下标当元素值用。该死的脑瓜
	// for v := range nums {
	// 	aMap[v] = aMap[v] + 1
	// }
	for k, c := range aMap {
		if c == 1 {
			return k
		}
	}
	return -1
}

func TestSingleNumber(t *testing.T) {
	r := singleNumber([]int{4, 1, 2, 1, 2})
	assert.Equal(t, 4, r)
}
