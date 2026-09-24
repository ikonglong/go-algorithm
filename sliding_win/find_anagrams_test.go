package slidingwin

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/

// 算法思路
// 1. 只创建滑动窗口的右侧指针 right（包含），左指针根据 right 窗口长度动态计算。
// 2. 使用 right 遍历 s。每次迭代表示将要处理 s[right]，而非已经处理完成
// 		初始时 Right 为 0
//      一旦 right 移动到第一个窗口之后，自此开始，right 每次移动，left 也要跟着移动
//      一旦第一个窗口完整了，自此开始，每次移动都要检查滑动窗口跟 target 是否是异位词

func findAnagrams(s string, p string) []int {
	if p == "" || len(s) < len(p) {
		return nil
	}

	result := make([]int, 0)
	// 假设输入只有 26 个小写字母
	var window [26]int
	var target [26]int
	winLen := len(p)

	for _, c := range p {
		target[c-'a']++
	}

	// 只创建滑动窗口的右侧指针 right，包含 right 边界；
	// 滑动窗口的左侧使用 right、winLen 动态计算。
	// 初始时，write 为 0，表示将要处理 s[0]，而非已经处理完成，
	// 因此在 for 循环之前没做什么初始化/准备工作
	for right := 0; right < len(s); right++ {
		window[s[right]-'a']++ // 先将当前字符加入 window

		// 一旦 right 移动到完整的第一个窗口之后，自此开始，right 每次移动，
		// left 也要跟着移动，这样窗口大小才保持不变
		if right >= winLen {
			left := right - winLen
			window[s[left]-'a']--
		}

		// 一旦第一个窗口完整了，自此开始，每次移动都要检查滑动窗口跟 target 是否是异位词
		if right >= winLen-1 && slices.Equal(window[:], target[:]) {
			result = append(result, right+1-winLen)
		}
	}

	return result
}

// 算法思路跟前一个相同，具体差异在「判断是否是异位词」。这个实现版本更简洁些
func findAnagrams_charDiff(s string, p string) []int {
	if p == "" || len(s) < len(p) {
		return nil
	}

	result := make([]int, 0)
	winLen := len(p)
	// 当前滑动窗口跟 P 之间的字符数量差值
	diff := make([]int, 26)
	// 当滑动窗口是 p 的异位词时，diff 的期望结果，用于比较判断
	zero := make([]int, 26)

	// 初始化 diff，将 p 中的字符数表示为负数
	for _, c := range p {
		diff[c-'a']--
	}

	for right := 0; right < len(s); right++ {
		diff[s[right]-'a']--

		// 自 right 移动到完整的第一个窗口之后开始
		if right >= winLen {
			left := right - winLen
			// 移动 left
			diff[s[left]-'a']--
		}

		if right >= winLen-1 && slices.Equal(diff, zero) {
			result = append(result, right-winLen+1)
		}
	}
	return result
}

// 算法思路：
// 1. 创建两个指针 i、j，初始时分别指向 0、m-1，m 为目标单词的长度。
//    i、j 分别是初始滑动窗口的左侧、右侧边界，闭区间。
// 2. 使用 j 遍历 s
//      每个迭代开始时，将要处理 s[j] 表示的滑动窗口
//      判断是否是意位词，若是，记录意位词开始位置
//      移动滑动窗口的左右边界 i、j

func findAnagrams2(s string, p string) []int {
	if p == "" || len(s) < len(p) {
		return nil
	}

	result := make([]int, 0, 8)
	pChars := []byte(p)
	slices.Sort(pChars)
	m := len(p)
	i, j := 0, m-1
	for j < len(s) {
		subStr := []byte(s[i : j+1])
		if isAnagram2(subStr, pChars) {
			// fmt.Printf("i:%d,j:%d  subStr:%s\n", i, j, subStr)
			result = append(result, i)
		}
		i++
		j++
	}
	return result
}

func isAnagram2(s []byte, t []byte) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [256]int
	// 挺巧妙的。如果 s, t 中每个字符的个数都相同，最终每个位置为 0
	for i := range s {
		counts[s[i]]++
		counts[t[i]]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram3(s []byte, t []byte) bool {
	if len(s) != len(t) {
		return false
	}
	slices.Sort(s)
	return slices.Equal(s, t)
}

func TestIsAnagram(t *testing.T) {
	assert.True(t, isAnagram2([]byte("bca"), []byte("abc")))
}
