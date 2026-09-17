package backtrack

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	result := make([]string, 0)

	// 数字按键
	digitToLetters := make(map[int][]byte)
	digitToLetters[2] = []byte{'a', 'b', 'c'}
	digitToLetters[3] = []byte{'d', 'e', 'f'}
	digitToLetters[4] = []byte{'g', 'h', 'i'}
	digitToLetters[5] = []byte{'j', 'k', 'l'}
	digitToLetters[6] = []byte{'m', 'n', 'o'}
	digitToLetters[7] = []byte{'p', 'q', 'r', 's'}
	digitToLetters[8] = []byte{'t', 'u', 'v'}
	digitToLetters[9] = []byte{'w', 'x', 'y', 'z'}

	// 数字字符串有几个字符，就有几个决策阶段
	// 选择列表：每个数字对应的字母列表。
	//   因为是计算组合，因此处理完成的数字，不会再回溯处理
	// 路径：已添加到组合中的字母
	var backtrack func(digits string, k int, comb []byte)
	backtrack = func(digits string, k int, comb []byte) {
		if k == len(digits) {
			result = append(result, string(comb))
			return
		}

		letters, _ := digitToLetters[int(digits[k]-'0')]
		for _, char := range letters {
			// 做出选择
			comb = append(comb, char)
			backtrack(digits, k+1, comb)
			// 恢复
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(digits, 0, []byte(nil))
	return result
}
