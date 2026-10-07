package dp

// https://leetcode.cn/problems/longest-common-subsequence/
// 一个字符串的子序列，是保持原有字符的相对顺序不变，删除或叫跳过某些字符后形成的新字符串
// LCS (Longest Common Subsequence)

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	// 状态转移表
	// 第一、二维下标分别是 text1、text2 从头开始的，每个阶段追加一个字符的子串的长度，
	// 也可以将它们看作 text1、text2 中字符的另一个索引，和原始字符串中的索引相比，这里从 1 开始。
	// 因此 dp[i][j] 表示 t1[0, i-1] 和 t2[0, j-1] 两个子串的最长公共子序列的长度
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// 初始化第 1 行
	for j := range m + 1 {
		// text1 长度为 0 的子串跟 text2 长度为 j 的子串之间最长公共子序列长度为零
		dp[0][j] = 0
	}
	// 初始化第 1 列
	for i := range n + 1 {
		// text1 长度为 i 的子串跟 text2 长度为 0 的子串之间最长公共子序列长度为零
		dp[i][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 从 dp 对字符的索引看，将要匹配 t1' 第 i 个字符和 t2' 第 j 个字符。
			// 但 dp 的索引从 1 开始，因此翻译为原始字符串的索引分别是 i-1 和 j-1
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
