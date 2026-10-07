package dp

// https://leetcode.cn/problems/partition-equal-subset-sum/description/
// 数组中全部是正整数

// 算法思路：
// - 分解问题：把给定数组分割成两个相等的子集，等价于选择部分元素放入一个集合，
//   使得集合元素和等于整个数组和的一半
// - 多阶段决策模型：
//   1. 数组有多少个元素，就有多少个决策阶段
//   2. 每个决策阶段决定当前元素是否放入集合
//   3. 路径就是状态转移表
// - 状态转移方程：
//   dp[i][j], 0<=i<n, 0<=j<=sum(nums[i])/2, 单元格是列j表示的和是否可达
//   dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 { // 奇数
		return false
	}
	// 创建状态转移表
	dp := make([][]bool, len(nums))
	for i, _ := range dp {
		// 第二维长度得是 sum+1，否则初始化时可能会下标越界。
		// 例如：[100,4,6]
		//dp[i] = make([]bool, sum/2+1)
		dp[i] = make([]bool, sum+1)
	}
	// 初始化第 0 行
	dp[0][0] = true
	dp[0][nums[0]] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= sum/2; j++ {
			// 不管走哪个分支，只要父节点是 dp[i-1][j]，就表示 nums[i] 不放入
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][sum/2]
}
