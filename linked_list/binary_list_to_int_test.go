package linkedlist

import (
	"fmt"
	"testing"
)

func getDecimalValue(head *ListNode) int {
	_, total := numForNode(head)
	return total
}

func numForNode(n *ListNode) (idx int, total int) {
	fmt.Printf("n:%+v\n", *n)
	if n.Next == nil {
		return 0, n.Val << 0
	}
	idx, total = numForNode(n.Next)
	return idx + 1, total + n.Val<<(idx+1)
}

func getDecimalValue2(head *ListNode) int {
	total := 0
	for node := head; node != nil; node = node.Next {
		// `total << 1`：将 `total` 的二进制整体左移一位，相当于乘以 2，并在末尾空出一个 `0`
		// `| node.Val`：通过按位或，把末尾设置为当前节点的值（`0` 或 `1`）
		total = total<<1 | node.Val
	}
	return total
}

func TestGetDecimalValue(t *testing.T) {
	testCases := []struct {
		head *ListNode
		want int
	}{
		{&ListNode{1, &ListNode{0, &ListNode{1, nil}}}, 5},
		{&ListNode{0, nil}, 0},
	}
	for _, tc := range testCases {
		got := getDecimalValue(tc.head)
		if got != tc.want {
			t.Errorf("getDecimalValue(%v) = %d, want: %d", tc.head, got, tc.want)
		}
	}
}
