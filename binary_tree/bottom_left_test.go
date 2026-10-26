package binary_tree

// https://leetcode.cn/problems/find-bottom-left-tree-value/
// 找出二叉树的 最底层 最左边 节点的值

// 算法思路：
//  1. 联想到按层遍历
//  2. 如果从遍历第一层的根节点开始，扩展单个节点总是从右向左，那么下一层节点
//     进入队列的顺序就是从右向左，进而每一层节点进入队列的顺序都是从右向左。
//     这样最后那个节点就是正确结果

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	result := -1
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		// 队首出队
		n := que[0]
		que = que[1:]
		result = n.Val
		// 扩展单个节点时，总是从右到左
		if n.Right != nil {
			que = append(que, n.Right)
		}
		if n.Left != nil {
			que = append(que, n.Left)
		}
	}
	return result
}
