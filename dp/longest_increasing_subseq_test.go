package dp

// https://leetcode.cn/problems/longest-increasing-subsequence/
// LIS (Longest Increasing Subsequence)

// 算法思路：
// dp[n], dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
// dp[i]=max(dp[j]+1)，dp[j]不是一个元素，而是所有满足
// 0<=j<i && nums[j]<nums[i] 的 nums[j] 对应的 dp[j]

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // 最短的递增子序列长度为 1，因此可以用 0 作为默认值
	dp[0] = 1            // 以 nums[0] 结尾的递增子序列仅包含它自身，因此长度为 1
	for i := 1; i < n; i++ {
		dp[i] = 1 // 这句非常容易忘记。任何一个元素自身就是一个递增子序列，长度为 1，应该作为默认值
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxLen := 0
	for _, v := range dp {
		maxLen = max(maxLen, v)
	}
	return maxLen
}

// minTails 数组就是经典的 patience sorting 里的 tails 数组（patience sorting 的牌堆顶），
// 满足 minTail[1] < minTail[2] < ... < minTail[lisLen]（严格递增是它能二分的前提）。

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	// 把此数组理解为映射：key 为子序列的长度，value 为该长度的所有子序列中尾元素的最小值
	// key 就是下标，子序列长度可为 n，因此长度为 n+1
	minTails := make([]int, n+1)
	end := 0 // seqLenToMinTail 数组中保存过最小值部分的上界（包含）

	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, n)
	// 下面这两行导致 nums[0] 子序列长度和最小值都没有记录到 lenToMinTail[1]，
	// 失败例子：[0,1,0,3,2,3]，处理 nums[1] 后将 lenToMinTail[1] 设置为 1 了，但应该是 2
	// dp[0] = 1 // 第0个元素自身就是一个以它结尾的最长子序列，因此长度为 1
	// for i := 1; i < n; i++ {
	for i := 0; i < n; i++ {
		lastIdx := lastIndexLessThan(minTails, end, nums[i])
		if lastIdx == -1 { // 未找到，表明前面没有比 nums[i] 小的元素
			dp[i] = 1
		} else {
			// nums[i] 并不一定接在以 nums[i-1] 结尾的子序列上
			// dp[i] = dp[i-1] + 1
			// lastIdx 就是尾元素最小值小于 nums[i] 的最后也即最长的子序列的长度
			dp[i] = lastIdx + 1
		}

		// 以 nums[i] 结尾的最长子序列长度计算出来了，需要更新对应长度尾元素最小值

		// 若以 nums[i] 结尾的最长子序列长度大于 end，直接保存尾元素最小值，更新 end
		if dp[i] > end {
			minTails[dp[i]] = nums[i]
			end = dp[i]
		} else if minTails[dp[i]] > nums[i] { // 若小于等于 end，还需检查长度对应的最小尾元素是否要更新
			minTails[dp[i]] = nums[i]
		}
	}

	return end

	// 遍历求最大值是多余的，end 就是最长长度
	// result := 0
	// for _, v := range dp {
	// 	result = max(result, v)
	// }
	// return result
}

// 从前一个版本看，DP 表也不需要了，改进后的版本
func lengthOfLIS3(nums []int) int {
	n := len(nums)
	// minTails[L] = 所有长度为 L 的递增子序列中，尾元素的最小值。
	// 有效下标 [1, lisLen]，且 minTails[1] < minTails[2] < ... < minTails[lisLen]（据此才能二分）
	minTails := make([]int, n+1)
	lisLen := 0 // seqLenToMinTail 数组中保存过最小值部分的上界（包含）

	for i := range n {
		lastIdx := lastIndexLessThan(minTails, lisLen, nums[i])
		len := 0
		if lastIdx == -1 { // 未找到，表明前面没有比 nums[i] 小的元素
			len = 1
		} else {
			// lastIdx 就是尾元素最小值小于 nums[i] 的最后也即最长的子序列的长度
			len = lastIdx + 1
		}

		// 以 nums[i] 结尾的最长子序列长度计算出来了，需要更新对应长度尾元素最小值

		// 若以 nums[i] 结尾的最长子序列长度大于 end，直接保存尾元素最小值，更新 end
		if len > lisLen {
			minTails[len] = nums[i]
			lisLen = len
		} else if minTails[len] > nums[i] { // 若小于等于 end，还需检查长度对应的最小尾元素是否要更新
			minTails[len] = nums[i]
		}
	}

	return lisLen
}

// lastIndexLessThan 在升序数组 a 的 [1, size] 区间内，
// 返回最后一个满足 a[i] < target 的下标；不存在返回 -1。
func lastIndexLessThan(a []int, size int, target int) int {
	low := 1
	high := size
	for low <= high {
		mid := (low + high) / 2
		if a[mid] < target { // a[mid] 都满足期望
			if mid == size || a[mid+1] >= target { // 找最后一个
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
