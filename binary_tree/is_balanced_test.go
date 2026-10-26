package binary_tree

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
// 判断一棵二叉树是否是平衡二叉树，无法直接采用递归思路，因为左右子树是平衡二叉树，并不能推导出原树也是平衡二叉树。
// 但可将原问题转化为求树的深度或者高度。在求高度的过程中，一旦检测到左右子树的高度相差超过 1，那么原树肯定不是平衡
// 二叉树，终止递归
//
// 注意，树的最大深度、树的深度、树的高度，这三种描述是等价的，不要被不同的描述迷惑

func isBalanced(root *TreeNode) bool {
	balanced := true

	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		// 剪枝触发点 1
		if !balanced {
			return 0
		}
		if root == nil {
			return 0
		}
		lHeight := height(root.Left)
		// 剪枝触发点 2：height 调用可能修改 balanced，因此调用返回后检查更自然。若为 false，更早剪枝，效果更好
		// if !balanced {
		// 	return 0
		// }
		rHeight := height(root.Right)
		// if !balanced {
		// 	return 0
		// }
		if absInt(lHeight-rHeight) > 1 {
			balanced = false
			// 走到这里，返回什么已经不重要了。0 虽然不是语义上正确的值，但不影响最终结果
			// 也可以这里不返回，继续执行，返回正确的值
			// return 0
		}
		return max(lHeight, rHeight) + 1
	}

	height(root)
	return balanced
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
