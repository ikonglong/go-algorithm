package hashmap

// https://leetcode.cn/problems/check-permutation-lcci/
// 判断一个字符串是不是由另一个字符串的字符重排得到的，其实就是异位词
// 算法思路：
// 1. 遍历第一个字符串，用一个 map 记录每个字符的出现次数
// 2. 遍历另一个字符串，每次将 map 中对应字符的计数减一
// 3. 遍历 Map，任何字符的计数不等于 0，就返回 false，否则返回 true
//
// 也可以用关联数组替代 map，即数组的下标表示字符的 ASCII 码，对应的元素值为字符计数

func CheckPermutation(s1 string, s2 string) bool {
	charFreq := make(map[byte]int)
	for _, c := range []byte(s1) {
		charFreq[c] = charFreq[c] + 1
	}
	for _, c := range []byte(s2) {
		// 这样写就埋了 bug："a", "ab"，只考虑了 s1,s2 都有的字符
		// if v, found := charFreq[c]; found {
		// 	charFreq[c] = v - 1
		// }
		charFreq[c] = charFreq[c] - 1
	}
	for _, v := range charFreq {
		if v != 0 {
			return false
		}
	}
	return true
}
