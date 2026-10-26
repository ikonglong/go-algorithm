package binary_tree

// https://leetcode.cn/problems/successor-lcci/
// 中序遍历中指定节点的后继节点

// 从实现版本 2 简化而来
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var prev, result *TreeNode
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if result != nil { // 剪枝，提前返回
			return
		}
		// 处理根节点
		if prev != nil {
			result = root
			return
		}
		if root == p {
			prev = p
		}
		inorder(root.Right)
		if result != nil { // 剪枝，提前返回
			return
		}
	}

	inorder(root)
	return result
}

func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var prev, result *TreeNode
	// 不需要，这里采用的不是循环，不需要循环变量
	// 采用递归，root 参数就相当于核心循环变量吧
	// curr := root
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if result != nil { // 剪枝，提前返回
			return
		}
		// 处理根节点
		// if prev == nil && root == p { // 正确
		if root == p {
			prev = p
		} else if prev != nil {
			result = root
			return
		}
		// 放这里埋了个 bug：如果 root 就是目标节点，逻辑是设置 prev 为 root后，
		// 立即判断 prev 是否为空？不为空，设置 result 为 root。root 不可能既是
		// 前驱节点，又是后继节点
		// if prev != nil {
		// 	result = root
		// 	return
		// }
		inorder(root.Right)
		if result != nil { // 剪枝，提前返回
			return
		}
		// 放这里也埋了个 bug：如果 root 就是目标节点，逻辑是设置 prev 为 root，
		// 遍历右子树，再判断 prev 是否为空？不为空，设置 result 为 root。同样，
		// root 不可能既是前驱节点，又是后继节点
		// if prev != nil {
		// 	result = root
		// 	return
		// }
	}

	inorder(root)
	return result
}

// 算法思路跟前面两个版本一样，只是实现细节上略有不同
func inorderSuccessor3(root *TreeNode, p *TreeNode) *TreeNode {
	coming := false // 只有找到目标节点，后继节点才算即将到来
	var successor *TreeNode
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if successor != nil {
			return
		}
		if coming {
			successor = root
			// 为什么找到后继节点后再重置为 false？
			// 因为找到时 coming 必定为 true。找到后如果将其重置为 false，它就不可能再变为 true 了，
			// 这样就可确保 successor 取到一个值后就不会再变了。
			//
			// 但我认为这行并非必需，因为不可能有两个节点的地址都等于 p，所以找到后当前递归树分支立即返回了；
			// 其他递归树分支会检测到 successor 不为空，然后也立即返回了。测试验证确实如此
			coming = false
			return
		}
		if root == p {
			coming = true
		}
		inorder(root.Right)
		// 这里不剪枝也正确。想想为什么？
		// 既然你说剪枝，就意味着已经确定解不在这些递归树分支了，因此剪与不剪都不影响最终结果
		if successor != nil {
			return
		}
	}

	inorder(root)
	return successor
}
