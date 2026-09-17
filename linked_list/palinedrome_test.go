package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(head *ListNode) bool {
	reversed := reverseAndClone(head)
	node, reversedNode := head, reversed
	for node != nil && reversedNode != nil {
		if node.Val != reversedNode.Val {
			return false
		}
		node = node.Next
		reversedNode = reversedNode.Next
	}
	return true
}

func reverseAndClone(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		newCurrent := &ListNode{
			Val: current.Val,
		}
		newCurrent.Next = prev
		prev = newCurrent
		current = next
	}
	return prev
}

// ======

//  1. 先将链表的前半部分反转。
//     如何找到链表中间位置？
//     根据链表长度是否已知，分两种情形：
//  1. 长度已知：就用一般 for 循环迭代到链表中间位置。
//  2. 长度未知：就用快慢指针，慢指针依次走一步，快指针依次走两步。当快指针走到链表末尾时，慢指针就走到了链表中间位置。
//     如何反转前半部分？链表反转很适合用栈：从头到尾遍历，逐个入栈，再逐个出栈即可反转。
//  2. 继续遍历链表的后半部分，并跟反转后的前半部分逐个比较。如果值相同就是回文链表，否则不是。
//
// Example:
// [1]
// [1, 2, 1]
// [1, 2, 2, 1]
// [1, 2, 1, 2, 1]
func isPalindrome2(head *ListNode) bool {
	fmt.Printf("input: %v\n", head.ValList())
	if head == nil || head.Next == nil {
		return true
	}

	// slow 指向后半段的第一个节点
	slow, fast := head, head
	count := 1
	fmt.Printf("[initial] slow: %v, fast: %v\n", slow, fast)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			count += 1
		} else {
			count += 2
		}
		fmt.Printf("[in loop] slow: %v, fast: %v\n", slow, fast)
	}
	fmt.Printf("len of list: %d\n", count)

	reversed := reverseAndClone2(head, slow)
	p1, p2 := slow, reversed
	// 初始时 fast 指针指向第一个节点，所以第一个节点是出发点，未被指针扫过，
	// 即第一个节点没有参与分组。因此如果有偶数个节点，尾节点所在分组只有它一个，
	// 这意味着 fast 最后指向 nil；如果有奇数个节点，fast 最后指向尾节点。
	if fast != nil {
		p1 = slow.Next
	}
	fmt.Printf("p1: %v, p2: %v\n", p1.ValList(), p2.ValList())
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1 == nil && p2 == nil
}

func reverseAndClone2(head *ListNode, dummyTail *ListNode) *ListNode {
	stack := make([]*ListNode, 0, 16)
	p := head
	// 入栈
	for p != dummyTail {
		node := &ListNode{
			Val: p.Val,
		}
		stack = append(stack, node)
		p = p.Next
	}
	fmt.Printf("stack: %v\n", stack)

	// 出栈，构建反转后的链表
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}

	return stack[len(stack)-1]
}

func isPalindrome2_Bad(head *ListNode) bool {
	fmt.Printf("input: %v\n", head.ValList())
	if head.Next == nil {
		return true
	}

	// fast 指向尾节点后的空位置，这导致当链表节点数为奇数时，slow 多跳了一次。
	// 例如：[1,2,1,2,1]，fast、slow 都跳三次，slow 指向索引为 3 的 2；
	// [1,2,2,1]，fast 原本就要跳两次，虽然最后一组只有一个节点，也需要跳这一次，
	// 所以对它来说，没有多跳一次。
	//
	// 这个选择原本期望的是，slow 总是停在后半部分的第一个节点，事实并非如此。
	// 要在这个思路上补救，只能在反转前半部分时，如果节点数是奇数，循环的结束条件是
	// p.next == slow，其他情形仍是 p == slow
	slow, fast := head, head
	stepCount := 0
	for fast != nil {
		fmt.Printf("[in loop] slow: %v, fast: %v\n", slow, fast)
		slow = slow.Next
		if fast.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
		stepCount++
	}
	fmt.Printf("slow:%v, fast:%v, stepCount:%d\n", slow, fast, stepCount)

	reversed := reverseAndClone2(head, slow)
	p1, p2 := reversed, slow
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1 == nil && p2 == nil
}

// ======

func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}
	palindrome, _ := isPalindromeRecurse(head, lenOf(head))
	return palindrome
}

// len 为像俄罗斯套娃一样将链表逐层切分为子链表，每个子链表的长度
// peer 为后半段的下一个节点，也是一次调用中跟前半段的 head 对应的节点
func isPalindromeRecurse(head *ListNode, len int) (palindrome bool, peer *ListNode) {
	if len == 0 { // 节点数为偶数 [1, 2, 2, 1]
		// 犯过的错误。对于 [1, 2, 2, 1]，len == 0 时，head 指向索引为 3 的 2
		// return true, nil
		return true, head
	} else if len == 1 { // 节点数为奇数
		return true, head.Next
	}

	palindrome, peer = isPalindromeRecurse(head.Next, len-2)

	// 如果发现已经不是回文，或者属于前半段的 head 没有可比的后半段节点（peer 为空）
	if !palindrome || peer == nil {
		return palindrome, peer
	}

	palindrome = (head.Val == peer.Val)
	peer = peer.Next
	return palindrome, peer
}

func TestReverse(t *testing.T) {
	linked1 := NewLinked([]int{1})
	assert.Equal(t, []int{1}, reverseAndClone(linked1).ValList())
	list2 := NewLinked([]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, reverseAndClone(list2).ValList())
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		head *ListNode
		want bool
	}{
		{NewLinked([]int{1, 2}), false},
		{NewLinked([]int{1, 2, 1}), true},
		{NewLinked([]int{1, 2, 2, 1}), true},
		{NewLinked([]int{1, 2, 1, 2, 1}), true},
	}
	for _, tc := range testCases {
		got := isPalindrome3(tc.head)
		if got != tc.want {
			t.Errorf("isPalindrome(%v)=%v, want: %v", tc.head.ValList(), got, tc.want)
		}
	}
}
