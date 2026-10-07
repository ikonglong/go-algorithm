package dp

import (
	"github.com/ikonglong/go-algorithm/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob3(root *common.TreeNode) int {
	result := postOrder(root)
	return max(result[0], result[1])
}

func postOrder(root *common.TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	// 左子树
	leftMoney := postOrder(root.Left)
	// 右子树
	rightMoney := postOrder(root.Right)
	// 当前节点
	return []int{
		// 不偷当前房间，左右子房间各有两个选择：偷、不偷
		max(leftMoney[0], leftMoney[1]) + max(rightMoney[0], rightMoney[1]),
		// 偷，左右子房间都只有一个选择：不偷
		leftMoney[0] + rightMoney[0] + root.Val,
	}
}
