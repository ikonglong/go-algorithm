package binary_tree

// https://leetcode.cn/problems/symmetric-tree/
// 算法思路：
// 求原树是否是镜像对称的，可分解或转化为求左子树和右子树是否是镜像对称的。
// 更进一步，求左子树和右子树是否是镜像对称，可分解为求这两个子树的左子树和右子树、
// 右子树和左子树分别是否是对称的。
// 若两对子树都是对称的，那么原树的左右子树就是对称的

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return twoSymmetric(root.Left, root.Right)
}

func twoSymmetric(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil {
		return false
	} else if root2 == nil {
		return false
	}

	// 这句容易漏
	if root1.Val != root2.Val {
		return false
	}
	lSymmetric := twoSymmetric(root1.Left, root2.Right)
	rSymmetric := twoSymmetric(root1.Right, root2.Left)
	return lSymmetric && rSymmetric
}
