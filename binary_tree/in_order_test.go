package binary_tree

import (
	"testing"
)

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorderTraverseFrom(root, &result)
	return result
}

func inorderTraverseFrom(root *TreeNode, result *[]int) {
	// root 就是当前节点
	if root == nil {
		return
	}

	inorderTraverseFrom(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraverseFrom(root.Right, result)
}

// 非递归实现思路：自己实现一个栈，模拟递归调用函数栈
// 模拟 压栈 / 扩展 / 出栈 的过程
// 1. 一直访问左，压栈。此时将栈帧 status 设为 1
// 2. 往左到底后， 考察栈顶节点
// 		a) 栈顶节点，若 status=1，变为 2，扩展右子节点。
//         扩展右子节点跟扩展左子节点的逻辑一样，因此再次从步骤 1 开始执行，继续循环
// 		b) 栈顶节点，若 status=2，变为 3，出栈。
//         处理左右子树期间压入的栈帧都弹出后，从步骤 2 开始执行，
// 		   这个时机可访问栈顶节点，根据需求做相应处理

func inorderTraverseLoop(root *TreeNode) []int { // leetcode name: inorderTraversal
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*CallFrame, 0)
	p := root // p 用来遍历树
	for {     // while true {}
		// topFrame := stack[len(stack)-1] // peek stack

		// 启动 p 左子树的处理。通过层层嵌套的 dfs(paramRoot.left) 一路沿左边向下递进，直到
		// root 树的最左子节点的 null 左孩子。注意，递进过程中没有启动任何节点右子树的处理。
		//
		// 每层调用中右子树的处理何时启动的呢？
		// 递进到底后，开始自底向上逐层回归。回归到某层调用中时，意味着参数节点的左子树已处理完成。
		// 接着继续执行此调用中的剩余逻辑，包括但不限于：打印参数节点值、处理右子树

		// 每个迭代开始时，栈顶帧 status 必定为 1，因此可以不检测
		// if stack[len(stack)-1].status == 1 {
		for p != nil {
			// 模拟 dfs(root.left) 调用以及执行期间发生的所有递归调用。将所有调用压栈
			stack = append(stack, &CallFrame{
				status: 1,
				node:   p,
			})
			p = p.Left // 一路向左递进
		}

		// 上面的循环沿左边一路递进到底后，这里通过弹出执行完成的栈帧实现自底向上逐层回归。
		// 若栈顶节点是其父节点的左子节点，则循环弹出左子树；若是右子节点，则循环弹出右子树和父节点。
		// 弹出后栈顶帧的 status 为 1。
		for len(stack) != 0 && stack[len(stack)-1].status == 2 {
			stack[len(stack)-1].status = 3
			stack = stack[:len(stack)-1] // 弹出栈帧，模拟函数调用返回
		}

		// 在弹栈后检测栈是否为空是自然直观合理的
		if len(stack) == 0 {
			break // 结束循环
		}

		// 虽然前面说 2 表示左子树已处理，但其实在这里表示的是左子树已处理，且右子树已开始处理或处理完成。
		// 一个节点的左子树处理完后，先将栈帧 status 设为 2，再开始处理右子树。处理右子树期间，肯定会有
		// 新的栈帧入栈，栈顶节点并不是此时的栈顶节点。右子树处理完后，走到第二部分的状态检测时栈顶帧
		// status 必定为 2，进而循环弹出右子树和此时的栈顶节点。因此，2 也表示左右子树都已处理。
		stack[len(stack)-1].status = 2
		// 埋有 bug：走到这里时的栈顶节点跟当前迭代开始时的栈顶节点很可能不同
		// result = append(result, topFrame.node.Val)
		result = append(result, stack[len(stack)-1].node.Val)
		// 将栈顶节点的右子树传给 dfs，下个迭代开始执行
		p = stack[len(stack)-1].node.Right
	}

	return result
}

// 中序遍历
func inorderTraversal_bad(root *TreeNode) []int {
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
func inorderTraversal_bad2(root *TreeNode) []int {
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
