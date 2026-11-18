package binary_tree

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

func pathTarget(root *TreeNode, target int) [][]int {
	result := make([][]int, 0)
	var dfs func(*TreeNode, int, []int, int)
	dfs = func(root *TreeNode, target int, path []int, pathSum int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		pathSum += root.Val
		// 如果走到叶子节点，则需判断是否是一个解
		if root.Left == nil && root.Right == nil {
			if pathSum == target {
				// 虽然 path 是按值传递，但后续处理可能会修改其底层数组，因此这里需要创建它的副本。
				// 埋有 bug：因为 dst slice 为空，导致什么也没拷贝
				// pathSnapshot := make([]int, 0)
				pathSnapshot := make([]int, len(path))
				copy(pathSnapshot, path)
				result = append(result, pathSnapshot)
			}

			// 回归前需要恢复状态。path 和 pathSum 都是当前调用栈帧中的局部变量，当前调用中的修改
			// 并不会影响外层调用中的同名变量，因此不做显式恢复回归后也会自然恢复。但 path 是 slice，
			// 这里先添加了值，回归前又通过重新切片恢复，不确定底层数组是否一起恢复了？不过底层数组是否
			// 一起恢复了并不重要，因为外层调用中向它的 path 添加值时会覆写底层数组中已有的值，若写入
			// 位置已有值的话
			// path = path[:len(path)-1]
			// 叶子节点也算是递归的最小子问题之一吧，因此这里需要返回
			// 从算法逻辑看，递进到叶子节点后，也应该回归或回溯
			return
		}

		dfs(root.Left, target, path, pathSum)
		dfs(root.Right, target, path, pathSum)
		// 执行结束自然返回前，也需要恢复状态
		// 因为 go slice 参数是按值传递，因此不做显式恢复回归后也自然恢复了。详细解释看上面
		// path = path[:len(path)-1]
	}

	dfs(root, target, []int(nil), 0)
	return result
}
