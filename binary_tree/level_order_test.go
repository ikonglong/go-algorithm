package binary_tree

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 层序遍历二叉树：输入：[8,17,21,18,null,null,6]，输出：[8,17,21,18,6]

// 算法思路：
// 1. 创建一个结果列表，用来保存遍历到的节点值
// 2. 创建一个列表，用于遍历
// 3. 初始时，先将 root 添加到队列
// 4. 重复如下步骤，直到队列为空
// 		a. 先将队首元素出队，再添加至结果列表。
// 		b. 展开队首元素，将其所有子节点添加至队列

func decorateRecord(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	que := append([]*TreeNode(nil), root)
	for len(que) != 0 {
		head := que[0]
		// 注意，对于 Go 语言，即使 que 长度为 1，
		// 创建切片时也可以这么写 que[1:]，不会报错
		que = que[1:]
		result = append(result, head.Val)
		if head.Left != nil {
			que = append(que, head.Left)
		}
		if head.Right != nil {
			que = append(que, head.Right)
		}
	}
	return result
}
