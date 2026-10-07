package dp

// https://leetcode.cn/problems/edit-distance/
// 编辑距离就是将两个单词变得一样所需的【最少操作数】
// 允许的操作：
// - 插入一个字符
// - 删除一个字符
// - 替换一个字符

func minDistance(word1 string, word2 string) int {
	w1Len := len(word1)
	w2Len := len(word2)

	// 状态转移表
	// 第一、二维下标分别是 word1、word2 从头开始的，每个阶段追加一个字符的子串的长度，
	// 也可以将它们看作 word1、word2 中字符的另一个索引，和原始字符串中的索引相比，这里从 1 开始。
	// 因此 dp[i][j] 表示 w1[0, i-1] 和 w2[0, j-1] 两个子串的编辑距离
	dp := make([][]int, w1Len+1)
	for i := range dp {
		dp[i] = make([]int, w2Len+1)
	}
	// 初始化第 1 列
	for i := range w1Len + 1 {
		// word1 长度为 i 的子串跟 word2 长度为 0 的子串之间的编辑距离是 i
		dp[i][0] = i
	}
	// 初始化第 1 行
	for j := range w2Len + 1 {
		// word1 长度为 0 的子串跟 word2 长度为 j 的子串之间的编辑距离是 j
		dp[0][j] = j
	}

	for i := 1; i <= w1Len; i++ {
		for j := 1; j <= w2Len; j++ {
			// 从 dp 对字符的索引看，将要匹配 w1' 第 i 个字符和 w2' 第 j 个字符。
			// 但 dp 的索引从 1 开始，因此翻译为原始字符串的索引分别是 i-1 和 j-1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[w1Len][w2Len]
}
