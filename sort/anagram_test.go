package sort_

import (
	"slices"
)

// https://leetcode.cn/problems/valid-anagram/
// 是否是字母异位词

// 先对两个字符串排序，然后逐个比较字符
func isAnagram(s string, t string) bool {
	// ok
	sChars := []byte(s)
	tChars := []byte(t)
	// ok
	// sort.Slice(sChars, func(i, j int) bool {
	// 	return sChars[i] < sChars[j]
	// })
	// sort.Slice(tChars, func(i, j int) bool {
	// 	return tChars[i] < tChars[j]
	// })
	slices.Sort(sChars)
	slices.Sort(tChars)
	return string(sChars) == string(tChars)
}

func isAnagram2(s string, t string) bool {
	sRunes := []rune(s)
	tRunes := []rune(t)
	slices.Sort(sRunes)
	slices.Sort(tRunes)
	return string(sRunes) == string(tRunes)
}

func isAnagram3(s string, t string) bool {
	sRunes := []rune(s)
	tRunes := []rune(t)
	slices.Sort(sRunes)
	slices.Sort(tRunes)
	return string(sRunes) == string(tRunes)
}

// 创建一个映射，键是字符，值是字符在字符串中出现的次数（字频）
// 先遍历字符串s，累计出每个字符的字频，再遍历字符串t，减去字符的字频
func isAnagram4(s string, t string) bool {
	sRunes := []rune(s)
	tRunes := []rune(t)
	rFreq := make(map[rune]int)
	for _, r := range sRunes {
		rFreq[r] = rFreq[r] + 1
	}
	for _, r := range tRunes {
		// 埋有 bug："a", "ab"
		// if v, found := rFreq[r]; found {
		// 	rFreq[r] = v - 1
		// }
		rFreq[r] = rFreq[r] - 1
	}
	for _, v := range rFreq {
		if v != 0 {
			return false
		}
	}
	return true
}
