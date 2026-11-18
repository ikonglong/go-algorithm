package binary_tree

// https://leetcode.cn/problems/path-sum-iii/

// 动态规划实现版本请看 dp/binarytree_pathsum_test.go

// 递归实现版本
func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	var dfs func(*TreeNode, int) map[int]int
	dfs = func(root *TreeNode, targetSum int) map[int]int {
		if root == nil {
			return map[int]int{}
		}
		leftMap := dfs(root.Left, targetSum)
		rightMap := dfs(root.Right, targetSum)

		// 处理当前节点

		// key 是路径节点累加和，value 是和对应的路径数
		retMap := make(map[int]int)
		// 这句非常关键。在回归过程中先将{当前节点值:1} 放入map，这样回归到后续节点时
		// 就可以计算以它为起点，以自身为终点的路径了。【注意】将当前节点放入map时，对
		// 应路径数总是 1，这样就可以正确处理一条路径上相邻节点值相同的情形
		retMap[root.Val] = 1
		for k, v := range leftMap {
			// 累加节点值，延伸路径至当前节点
			newK := k + root.Val
			// 延伸至当前节点的路径条数等于延伸至上一个节点的路径数
			newV := v
			// bug：添加这个检测的意图是如果 newK 在 leftMap 中存在，那么
			// 需要把对应的路径数加上。但弄巧成拙，这里加一次，遍历到它时又加
			// 一次。此外，若不加 newK != k，newK == k 时会把当前和对应的
			// 路径数加两次
			// if oldV, found := leftMap[newK]; newK != k && found {
			// 	newV += oldV
			// }
			if oldV, found := retMap[newK]; found {
				newV += oldV
			}
			retMap[newK] = newV
		}
		for k, v := range rightMap {
			// 累加节点值，延伸路径至当前节点
			newK := k + root.Val
			// 延伸至当前节点的路径条数等于延伸至上一个节点的路径数
			newV := v
			// bug：添加这个检测的意图是如果 newK 在 rightMap 中存在，那么
			// 需要把对应的路径数加上。但弄巧成拙，这里加一次，遍历到它时又加
			// 一次。此外，若不加 newK != k，newK == k 时会把当前和对应的
			// 路径数加两次
			// if oldV, found := rightMap[newK]; newK != k && found {
			// 	newV += oldV
			// }
			if oldV, found := retMap[newK]; found {
				newV += oldV
			}
			retMap[newK] = newV
		}
		for k, v := range retMap {
			if k == targetSum {
				count += v
				break
			}
		}
		return retMap
	}

	dfs(root, targetSum)
	return count
}
