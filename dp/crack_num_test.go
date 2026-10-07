package dp

import "strconv"

// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

// dp[n+1], dp[i] 字符长度为 i 的数有多少种解密方式或结果
// dp[i] = dp[i-1] + dp[i-2]
// dp[i-1]、dp[i-2] 分别为：字符长度为 i 时，最后一步可解码一个字符、两个字符，分别
// 对应前一步状态 dp[i-1], dp[i-2]
//
// 注意，两位数时可能超过 25，可能小于 10

func crackNumber(ciphertext int) int {
	s := strconv.Itoa(ciphertext)
	n := len(s)
	dp := make([]int, n+1)
	// 在这道题里 dp[0]=1 的理由和砍竹子那道题（cut_bamboo_test.go）不一样：
	// 这里是计数，计数问题里「空的前缀」天然算作 1 种方案。
	//
	// 最直观的理解：
	// dp[i] 表示「前 i 个字符有多少种解码方式」，那么 dp[0] 就是「前 0 个字符（空串）有多少种解码方式」。
	// 答案显然不是 0：空串本身就已经是「解好码」的状态了，你什么都不用做，这恰好是一种解法。所以 dp[0] = 1。
	// 0 的含义是「没有任何一种解法」——可「什么都不译」明明是一种译法，所以只能是 1。
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1]
		if isValid([]byte(s), i-2, i-1) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func isValid(chars []byte, start, end int) bool {
	num := (chars[start]-'0')*10 + (chars[end] - '0')
	return num >= 10 && num <= 25
}
