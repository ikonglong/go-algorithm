package binary_tree

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

// 这个实现的心智模型：采用深度优先遍历，先一路递进到底（叶子节点的空子节点），
// 然后开始回归，回归过程中，开始产生路径，路径自底向上延伸。左、右子树调用回归
// 后，计算以 root 为峰顶节点的最大路径和（root.val + 从左子树底部向上延伸到
// 左子树根的最大路径和 + 从右子树底部向上延伸到右子树根的最大路径和）。然后返
// 回从当前树底部（即左、右子树底部）向上延伸到当前节点的最大路径和给父节点树调
// 用，帮助它计算以父节点为峰顶节点的最大路径和
func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var dfs func(*TreeNode) int
	// 返回值是自树底向上延伸到 root 的路径中最大路径和
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMaxPath := dfs(root.Left)
		rightMaxPath := dfs(root.Right)

		// 计算以当前节点为峰顶节点的路径的最大和
		// 以当前节点为峰顶节点的路径形状呈倒 V 型。它有两类变体：
		// 1. 只有一条手臂
		// 2. 没有双臂，只有峰顶节点
		max := root.Val
		if leftMaxPath > 0 {
			max += leftMaxPath
		}
		if rightMaxPath > 0 {
			max += rightMaxPath
		}
		if max > result {
			result = max
		}

		// 如果路径要穿过当前节点 root 继续向上延伸，那么需要计算位于
		// 当前树中的多条路径的最大值，有三种走法：走到 root 就结束、向左、向右。
		// 注意，不能先向左，再返回 root，然后向右，这样的路径包含 root 两次，
		// 不符合题目要求

		// 计算自当前树底层节点向上延伸至当前节点路径的最大和
		// 下面这行代码很重要，让 root 作为一个路径参与比较，使得
		// 不以叶子节点开始，不以原树根节点结束的路径有机会参与比较
		ret := root.Val
		if ret < root.Val+leftMaxPath {
			ret = root.Val + leftMaxPath
		}
		if ret < root.Val+rightMaxPath {
			ret = root.Val + rightMaxPath
		}
		return ret
	}
	dfs(root)
	return result
}

// 这个实现版本中 leftMax = dfs(root.Left) 实际算的是"左子树里随便哪条路径的最大值"，
// 这条路径可能在左子树深处、根本不经过 root.Left，于是 leftMax + root.Val 就把两段
// 不相连的东西加在一起了——这不是一条路径。这就是为什么你怎么补条件都会漏：原料本身就不是
// 能拼接的形状。
func maxPathSum_bug2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) int
	// 拿[-3]走查验证下逻辑
	dfs = func(root *TreeNode) int {
		// if root == nil {
		// 	return math.MinInt
		// }
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		maxSum := math.MinInt
		if root.Left != nil {
			leftMax := dfs(root.Left)
			maxSum = max(maxSum, leftMax)
			maxSum = max(maxSum, leftMax+root.Val)
		}
		if root.Right != nil {
			rightMax := dfs(root.Right)
			maxSum = max(maxSum, rightMax)
			maxSum = max(maxSum, rightMax+root.Val)
			maxSum = max(maxSum, maxSum+rightMax)
		}
		return maxSum
	}

	return dfs(root)
}

func maxPathSum_bug1(root *TreeNode) int {
	var dfs func(*TreeNode) int
	// 拿[-3]走查验证下逻辑
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
			// return math.MinInt
		}
		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		maxSum := max(leftMax, rightMax)
		maxSum = max(maxSum, leftMax+rightMax+root.Val)
		return maxSum
	}
	return dfs(root)
}
