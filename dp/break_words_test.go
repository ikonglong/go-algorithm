package dp

// https://leetcode.cn/problems/word-break/
// 注意小争哥的 PPT 中题目描述是按空格拆分，但 LeetCode 题目并没有说按空格拆分

// dp[n+1], dp[i] 长度为 i 的字符串是否可以拆分为给定单词列表中的单词
// dp[i] = dp[i-len(words[0])] || dp[i-len(words[1])] || dp[i-len(words[...])] || dp[i-len(words[n-1])]
//
// 套用爬楼梯模型分析这道题：
// 每个字符对应一个处理阶段
// 每个阶段，每一步爬几个台阶？
// 初看未明确，仔细分析，步长就是单词列表中每个单词的长度。单步走法等于单词唯一长度个数
// 这一点跟零钱兑换有多种面值可选挺类似，请看 coin_change_test.go 中基于爬楼梯模型的实现版本

func wordBreak(s string, wordDict []string) bool {
	wordN := len(wordDict)
	wordMap := make(map[string]struct{}, wordN)
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	n := len(s)
	dp := make([]bool, n+1)
	// 空前缀字符串被认为可以由单词列表中的零个单词拼接出来，因此值为真
	// 假设 s="hello"，验证下是否 ok
	dp[0] = true
	for i := 1; i <= n; i++ {
		for k, _ := range wordMap {
			// 举例确定切片，否则很容易出错。s="catdog" dp看作" catdog"
			if i >= len(k) && k == s[i-len(k):i] && dp[i-len(k)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
