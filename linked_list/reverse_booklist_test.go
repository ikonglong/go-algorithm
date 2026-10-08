package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBookList(head *ListNode) []int {
	result := make([]int, 0)
	var reverse func(*ListNode)
	// 用闭包直接捕获外层的 result 变量
	reverse = func(head *ListNode) {
		if head == nil {
			return
		}
		reverse(head.Next)
		result = append(result, head.Val)
	}
	reverse(head)
	return result
}

func reverseBookList2(head *ListNode) []int {
	result := make([]int, 0)
	var reverse func(*ListNode, *[]int)
	reverse = func(head *ListNode, result *[]int) {
		if head == nil {
			return
		}
		reverse(head.Next, result)
		// 等号左边 * 号：取出 result 指针指向的那个变量
		*result = append(*result, head.Val) // 解引用后写回
	}
	reverse(head, &result)
	return result
}

func reverseBookList3(head *ListNode) []int {
	var reverse func(*ListNode) []int
	reverse = func(head *ListNode) []int {
		if head == nil {
			return nil
		}
		result := reverse(head.Next)
		return append(result, head.Val)
	}
	return reverse(head)
}

// ```
//   外层 result ──► header{ arr, len=0, cap=0 }
//
//   调用 reverse(head, result)
//                        │ 拷贝 header
//                        ▼
//   内层 result ──► header{ arr, len=0, cap=0 }   (独立副本)
//       append 后    header{ arr', len=1, ... }   (改了副本，外层看不到)
// ```
//
// │ 补充：即使 make([]int, 0, n) 预分配了容量，append 会往同一个底层数组写数据，但长度 len 仍在外层副本里没
// │ 变，所以外层 result 依然是空的。长度才是决定"看到几个元素"的关键。
//
// 为什么内层的"更深一层"也传不回来
//
// 这是同一个问题的递归放大：每一层都拿到上一层的 result 副本，各自 append 各自改，互相都看不到，整条链上的
// append 全部丢失。

func reverseBookList_bug(head *ListNode) []int {
	result := make([]int, 0)
	var reverse func(*ListNode, []int)
	reverse = func(head *ListNode, result []int) {
		if head == nil {
			return
		}
		reverse(head.Next, result)
		result = append(result, head.Val)
	}
	reverse(head, result)
	return result
}
