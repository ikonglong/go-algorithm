package dp

// https://leetcode.cn/problems/house-robber-ii/description/
// 题目分析：最后一间房屋跟第一间房屋相邻，不能同时偷。
// 例如，有三间房屋：2 - 3 - 2
// 排成一         偷  不偷  偷，ok
// 围成一圈       偷  不偷  偷，not ok
// 可把房屋之间的连线理解为禁止同时偷的约束。
// 单排： A —— B —— C
// 成环： A —— B —— C
//        └─────────┘
//
// 若最后一间n-1不偷，可偷范围为 [0, n-2]。若最后一间偷，那么第一间不可偷，
// 可偷范围为 [1, n-1]，注意不是 [1, n-3]。虽然最后一间偷，则 n-2 不可偷，
// 但最后一间n-1可偷。不包含n-1就没把其价值累加上，结果就是错误的。

func rob2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	return max(robRange(&nums, 0, n-2), robRange(&nums, 1, n-1))
}

func robRange(nums *[]int, start, end int) int {
	numsV := *nums
	dp := make([][]int, len(numsV))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[start][0] = 0
	dp[start][1] = numsV[start]

	for i := start + 1; i <= end; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + numsV[i]
	}
	return max(dp[end][0], dp[end][1])
}
