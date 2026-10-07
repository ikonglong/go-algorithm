package binary_tree

import (
	"github.com/ikonglong/go-algorithm/common"
)

// 跟 binary_tree/node.go 中 ValList 实现相比较，这个实现优雅多了
func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 用一个队列来保存每一层的节点。处理树的一层时，逐个从队头取出并移除节点，
	// 再将节点值添加到该层的值列表中，同时将下一层节点添加到队列中。
	// 注意，每次迭代开始时队列的大小即为这一层的节点数量。
	result := make([][]int, 0)
	nodeQue := make([]*common.TreeNode, 0)
	nodeQue = append(nodeQue, root)
	for len(nodeQue) > 0 { // 队列中没有节点了，表明已遍历完所有的树节点
		levelLen := len(nodeQue)
		level := make([]int, 0, levelLen)
		for i := 1; i <= levelLen; i++ {
			n := nodeQue[0]
			level = append(level, n.Val) // 取队头元素
			nodeQue = nodeQue[1:]        // 移除队头元素
			// 将下一层节点添加到队列中
			if n.Left != nil {
				nodeQue = append(nodeQue, n.Left)
			}
			if n.Right != nil {
				nodeQue = append(nodeQue, n.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
