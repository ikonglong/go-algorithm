package binary_tree

import (
	"testing"

	"github.com/ikonglong/go-algorithm/common"
)

func inorderTraversal(root *common.TreeNode) []int {
	result := make([]int, 0)
	inorderTraversalRecurse(root, &result)
	return result
}

func inorderTraversalRecurse(root *common.TreeNode, result *[]int) {
	// root 就是当前节点
	if root == nil {
		return
	}

	inorderTraversalRecurse(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversalRecurse(root.Right, result)
}

// 中序遍历
func inorderTraversal_bad(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// 将叶子节点作为基线条件
	if root.Left == nil && root.Right == nil {
		list := make([]int, 0, 1)
		list = append(list, root.Val)
		return list
	}

	// 对于叶子节点，逻辑上要先执行上面的基线分支，再执行下面的通用逻辑。但是，执行基线
	// 分支就直接返回包含其值的列表了，不会再执行下面的通用逻辑。然而，下面的通用逻辑却
	// 依然包含对叶子节点的处理，即如果是叶子节点，就不再递归调用。因此，叶子节点的处理
	// 逻辑就重复实现了。

	// 对于叶子节点：root.Left == nil，root.Right == nil，执行下面的通用逻辑
	// ，最终仍是自然返回 []int{root.Val}，这和前面的基线返回的结果完全相同。
	// 因此可以删除叶子节点基线分支，保留子节点判断，就得到 inorderTraversal_bad2

	// 在已经判断「子节点非空才递归」的前提下，单独处理叶子节点的基线分支是多余的。
	var left, right []int
	if root.Left != nil {
		left = inorderTraversal_bad(root.Left)
	}
	if root.Right != nil {
		right = inorderTraversal_bad(root.Right)
	}

	list := make([]int, 0, len(left)+len(right)+1)
	list = append(list, left...)
	list = append(list, root.Val)
	list = append(list, right...)
	return list
}

// 这个写法递归何时开始返回的，而不会无限递归？
// 它在到达叶子节点时开始返回。虽然没有显式写出 `return` 形式的基线条件，
// 但两个 `if` 共同构成了递归终止条件。
// 例如树是：
//
//	  2
//	 / \
//	1   3
//
// 调用过程如下：
// inorderTraversal(2)
// ├── inorderTraversal(1)
// │   ├── Left == nil，不递归
// │   ├── Right == nil，不递归
// │   └── 返回 [1]
// ├── inorderTraversal(3)
// │   ├── Left == nil，不递归
// │   ├── Right == nil，不递归
// │   └── 返回 [3]
// └── 组合并返回 [1, 2, 3]
//
// 因此，递归的基线不一定非要写成：
//
//	if 条件 {
//	  return
//	}
func inorderTraversal_bad2(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var left, right []int

	if root.Left != nil {
		left = inorderTraversal_bad2(root.Left)
	}
	if root.Right != nil {
		right = inorderTraversal_bad2(root.Right)
	}

	list := make([]int, 0, len(left)+len(right)+1)
	list = append(list, left...)
	list = append(list, root.Val)
	list = append(list, right...)
	return list
}

func TestInorderTraversal(t *testing.T) {

}
