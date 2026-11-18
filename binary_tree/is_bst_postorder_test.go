package binary_tree

import "math"

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
// 判断给定数组是否为二叉搜索树的后序遍历结果。
// 注意，输入可能不是 BST 的后序结果，遍历过程中还要判断是否是 BST

func verifyTreeOrder(postorder []int) bool {
	// 定义全局结果变量的目的是剪枝，但发现对于这道题，不需要用它帮助实现剪枝
	// isValid := true
	var isBSTPostorderRange func([]int, int, int) bool
	// 拿单节点树走查验证下逻辑，确认 start == end 是否需要单独处理
	isBSTPostorderRange = func(post []int, start, end int) bool {
		if start > end {
			return true
		}

		// 分割出左右子树区间。同时这也确保了左子树所有节点都小于 root
		k := start
		// for k < end && post[k] < post[end] {
		for post[k] < post[end] {
			k++
		}

		// 检查右子树所有节点是否都大于 root
		for i := k; i < end; i++ {
			if post[i] < post[end] {
				// isValid = false
				return false
			}
		}

		leftValid := isBSTPostorderRange(post, start, k-1)
		if !leftValid {
			return false
		}
		rightValid := isBSTPostorderRange(post, k, end-1)
		return rightValid
		// 很啰嗦的写法
		// if !rightValid {
		// 	return false
		// }
		// return true
	}

	return isBSTPostorderRange(postorder, 0, len(postorder)-1)
}

// 用全局变量保存结果，以实现剪枝
func verifyTreeOrder2(postorder []int) bool {
	if len(postorder) == 0 {
		// return false
		return true
	}
	isBSTPostorder := true
	var isBSTPostorderRange func([]int, int, int) bool
	// 拿单节点树走查验证下逻辑，确认 start == end 是否需要单独处理
	isBSTPostorderRange = func(post []int, start, end int) bool {
		if start > end {
			return true
		}

		// 后序序列，最后一个元素就是当前树根节点值
		rootV := post[end]
		// 从右向左遍历数组，找到第一个小于 root 的，记为 k。
		// [start, k]、[k+1, end-1] 分别是左右子树区间
		//
		// 其实，划分出右子树的过程同时确保了右子树所有节点都大于根节点，
		// 下面就不用再检查右子树所有节点是否都大于根节点了
		k := end
		for k >= start && post[k] > rootV {
			k--
		}

		// 不能从左、右子树是 BST 后序序列推导出整个序列也是 BST 后序序列
		// 检查左子树所有节点是否都小于 root，右子树所有节点是否都大于 root
		for i := start; i <= k; i++ {
			if post[i] > rootV {
				isBSTPostorder = false
				return isBSTPostorder
			}
		}
		for i := k + 1; i < end; i++ {
			if post[i] < rootV {
				isBSTPostorder = false
				return isBSTPostorder
			}
		}

		// leftIsBSTPostorder or leftValid
		leftValid := isBSTPostorderRange(post, start, k)
		if !isBSTPostorder { // 剪枝
			return isBSTPostorder
		}
		rightValid := isBSTPostorderRange(post, k+1, end-1)
		if !isBSTPostorder { // 剪枝
			return isBSTPostorder
		}

		isBSTPostorder = leftValid && rightValid
		return isBSTPostorder
	}

	isBSTPostorderRange(postorder, 0, len(postorder)-1)
	return isBSTPostorder
}

// 没有剪枝的实现版本
func verifyTreeOrder3(postorder []int) bool {
	if len(postorder) == 0 {
		// return false
		return true
	}
	return isBSTPostorderRange3(postorder, 0, len(postorder)-1)
}

// 拿单节点树走查验证下逻辑，确认 start == end 是否需要单独处理
func isBSTPostorderRange3(post []int, start, end int) bool {
	if start > end { // 空区间
		return true
	}

	// 后序序列，最后一个元素就是当前树根节点值
	rootV := post[end]
	// 从右向左遍历数组，找到第一个小于 root 的，记为 k。
	// [start, k]、[k+1, end-1] 分别是左右子树区间
	k := end
	for k >= start && post[k] >= rootV {
		k--
	}

	// 不能从左、右子树是 BST 后序序列推导出整个序列也是 BST 后序序列
	// 检查左子树所有节点是否都小于 root，右子树所有节点是否都大于 root
	for i := start; i <= k; i++ {
		if post[i] > rootV {
			return false
		}
	}
	for i := k + 1; i < end; i++ {
		if post[i] < rootV {
			return false
		}
	}

	isLeftBSTPost := isBSTPostorderRange3(post, start, k)
	isRightBSTPost := isBSTPostorderRange3(post, k+1, end-1)
	return isLeftBSTPost && isRightBSTPost
}

// 带剪枝，且通过返回当前树最大值去掉前一个版本中遍历左右子树的两个 for 循环
// 算法思路有逻辑错误：递归调用左子树时，应该返回最大值；递归调用右子树时，应该返回最小值。
// 但编写时脑子里想的却是都返回最大值。但是力扣的测试居然通过了，应该是没有对应的测试用例触发 bug
func verifyTreeOrder4_bug(postorder []int) bool {
	if len(postorder) == 0 {
		// return false
		return true
	}
	isBSTPostorder := true
	var isBSTPostorderRange func([]int, int, int) (bool, int)
	// 拿单节点树走查验证下逻辑，确认 start == end 是否需要单独处理
	// 返回结果：是否是 BST && postorder；最大值
	isBSTPostorderRange = func(post []int, start, end int) (bool, int) {
		if start > end {
			return true, math.MinInt
		}

		k := end
		rootV := post[end]
		max := post[end]
		for k >= start && post[k] >= rootV {
			max = post[k]
			k--
		}

		leftValid, leftMax := isBSTPostorderRange(post, start, k)
		if !isBSTPostorder {
			return isBSTPostorder, max
		}
		rightValid, rightMax := isBSTPostorderRange(post, k+1, end-1)
		if !isBSTPostorder {
			return isBSTPostorder, max
		}
		// bug: 没有考虑 leftMax 和 rightMax 可能为 math.MinInt，表示空区间
		// isBSTPostorder = leftValid && rightValid && leftMax < rootV && rightMax > rootV
		isBSTPostorder = leftValid && rightValid &&
			(leftMax == math.MinInt || leftMax < rootV) &&
			(rightMax == math.MinInt || rightMax > rootV)
		return isBSTPostorder, max
	}

	isBSTPostorderRange(postorder, 0, len(postorder)-1)
	return isBSTPostorder
}
