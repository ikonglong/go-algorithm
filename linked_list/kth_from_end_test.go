package linkedlist

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

// 求倒数第 k 个节点，转化为求正数第 n-k+1 个节点
// leetcode name: trainingPlan
func kthNodeFromEnd(head *ListNode, cnt int) *ListNode {
	n := LenOf(head)
	kFromStart := n - cnt + 1
	count := 0
	p := head
	for p != nil {
		count++
		if count == kFromStart {
			return p
		}
		p = p.Next
	}
	return nil
}

// 借助递归思想，先递进到链表尾，在归的过程中计数，找到倒数第 k 个
// leetcode name: trainingPlan
func kthNodeFromEnd2(head *ListNode, cnt int) *ListNode {
	count := 0
	var result *ListNode
	var traverse func(*ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}

		traverse(head.Next)
		if result != nil { // 剪枝
			return
		}
		count++
		if count == cnt {
			result = head
			return
		}
	}
	traverse(head)
	return result
}
