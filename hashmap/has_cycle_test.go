package hashmap

// https://leetcode.cn/problems/linked-list-cycle/
// 题意分析：
// 1. 若有环，则一定是后面的节点链接回前面的节点，不可能是前面的节点通过某个连接链接至后面的节点（非下一个）
//
// 其他解法，看 week2/list_has_cycle.go

func hasCycle(head *ListNode) bool {
	aMap := make(map[*ListNode]struct{})
	p := head
	for p != nil {
		if _, found := aMap[p]; found {
			return true
		} else {
			aMap[p] = struct{}{}
		}
		p = p.Next
	}
	return false
}
