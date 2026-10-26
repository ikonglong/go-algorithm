package binary_tree

import "github.com/ikonglong/go-algorithm/common"

type TreeNode = common.TreeNode

// 函数调用栈帧
type CallFrame struct {
	// 函数调用内部状态：
	// 1: 左右子树都未处理
	// 2: 左子树已处理完，右子树还未处理
	// 3: 左右子树都已处理完
	status int
	// 模拟递归实现中的函数调用
	node *TreeNode
}
