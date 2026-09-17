package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(levelOrderList []any) *TreeNode {
	if len(levelOrderList) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: levelOrderList[0].(int),
	}

	parents := []*TreeNode{root}
	// level := 1 // root 是第 0 层
	for i := 1; i < len(levelOrderList); {
		// next_parents := make([]*TreeNode, 0, 1<<level)
		next_parents := make([]*TreeNode, 0, len(parents)*2)
		for _, p := range parents {
			if v, ok := levelOrderList[i].(int); ok {
				p.Left = &TreeNode{Val: v}
				next_parents = append(next_parents, p.Left)
			}
			i++
			if v, ok := levelOrderList[i].(int); ok {
				p.Right = &TreeNode{Val: v}
				next_parents = append(next_parents, p.Right)
			}
			i++
		}
		parents = next_parents
		// level++
	}

	return root
}

// 跟 binary_tree/level_order_test.go 实现相比较，这个实现挺丑陋了
func ValList(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	vList := make([]any, 0)
	// levelNodes 就是循环变量
	levelNodes := []*TreeNode{root}
	for {
		// 改进了实现，就不需要下面这段了
		// levelNull := true
		// for _, n := range levelNodes {
		// 	if n != nil {
		// 		levelNull = false
		// 		break
		// 	}
		// }
		// if levelNull {
		// 	break
		// }

		nextLevelNodes := make([]*TreeNode, 0)
		// 是否下一层所有节点都为空
		nextLevelNull := true
		for _, n := range levelNodes {
			if n == nil {
				vList = append(vList, "null")
			} else {
				vList = append(vList, n.Val)
				// 这里不能这样做：判断 left 和 right 是否为空，为空就不添加。
				// 因为还不能确定跟他们同一层的其他节点是否都为空。只有都为空，
				// 才没有必要添加到列表中。因为遍历时一次只能查看一个节点，无法
				// 一次查看同一层的所有节点，所以这里选择先将它们都添加到列表中，
				// 在下一次迭代时检查是否都为空
				nextLevelNodes = append(nextLevelNodes, n.Left, n.Right)
				if n.Left != nil || n.Right != nil {
					nextLevelNull = false
				}
			}
		}
		if nextLevelNull { // 下一层所有节点都为空
			break
		} else {
			levelNodes = nextLevelNodes
		}
	}

	return vList
}

func ValList2(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}

	vList := make([]any, 0)
	traverseByLevel([]*TreeNode{root}, &vList)
	return vList
}

func traverseByLevel(levelNodes []*TreeNode, vList *[]any) {
	if len(levelNodes) == 0 {
		return
	}

	nextLevelNodes := make([]*TreeNode, 0)
	// 为了省略叶节点的下一层，因为叶节点的左右子节点全为空，
	// 所以这一层全为空，可以安全省略
	nextLevelNull := true
	for _, n := range levelNodes {
		if n == nil {
			*vList = append(*vList, "null")
		} else {
			*vList = append(*vList, n.Val)
			nextLevelNodes = append(nextLevelNodes, n.Left, n.Right)
			if n.Left != nil || n.Right != nil {
				nextLevelNull = false
			}
		}
	}
	if nextLevelNull {
		nextLevelNodes = nil
	}

	traverseByLevel(nextLevelNodes, vList)
}
