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
func pathSum(root *common.TreeNode, targetSum int) int {
	count := 0
	var dfs func(*common.TreeNode, int) map[int]int
	dfs = func(root *common.TreeNode, targetSum int) map[int]int {
		if root == nil {
			return nil
		}

		// key 是以当前节点为上端点的路径和，value 是以当前节点为上端点且路径和为 ${key} 的路径条数
		rootResult := make(map[int]int)
		leftResult := dfs(root.Left, targetSum)
		rightResult := dfs(root.Right, targetSum)
		// 当前节点自身也独立组成一个路径，先处理该路径
		rootResult[root.Val] = 1
		for leftK, leftV := range leftResult {
			k := leftK + root.Val
			rootResult[k] = rootResult[k] + leftV

			// 埋有 bug，给 left child 添加前缀 root 后和为 k 的路径可能不止一条，这里只加了 1
			// rootResult[k] = rootResult[k] + 1

			// 埋有 bug：若 k 不存在，v 应该为 1，实际为 0
			// v, found := result[k]
			// if found {
			// 	v = v + 1
			// }
			// result[k] = v
		}
		for rightK, rightV := range rightResult {
			k := rightK + root.Val
			rootResult[k] = rootResult[k] + rightV
		}
		count += rootResult[targetSum]
		return rootResult
	}

	dfs(root, targetSum)
	return count
}
