package binary_tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 求一棵树第 k 大节点，无法直接采用递归思路解决。因为左、右子树第 k 大节点，不一定是原树第 k 大节点；
// 也无法基于左、右子树第 k_left, k_right 大节点，推导出原树第 k 大节点。
// 但可将原问题 转化为 按“右-根-左”中序遍历。遍历的同时，数节点个数，数到第 k 个，它就是第 k 大节点

func findTargetNode(root *TreeNode, k int) int {
	count := 0  // 遍历的同时，数节点个数
	result := 0 // 第k大节点的值

	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		// 剪枝触发点 1
		if count >= k {
			return
		}
		if root == nil {
			return
		}
		inorder(root.Right)
		// 处理根节点。若是真的遍历，这里要将根节点值添加到结果列表
		count++ // 把根节点数上
		// 剪枝触发点 2：递归调用可能修改 count，因此调用返回后检查更自然。若找到，更早剪枝，效果更好
		if count == k {
			result = root.Val
			return // 找到了，终止递归
		}
		inorder(root.Left)
		// 剪枝触发点 2：递归调用可能修改 count，因此调用返回后检查更自然。若找到，更早剪枝，效果更好
		// if count >= k {
		// 	return
		// }
	}

	inorder(root)
	return result
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Right)
	// 若是真的遍历，这里要将根节点值添加到结果列表
	// root.Val
	inorder(root.Left)
}
