package stack

import (
	"fmt"
	"testing"
)

type StackOfPlates struct {
	stacks []*SliceStack[int]
	cap    int
}

// leetcode name: Constructor
func NewStackOfPlates(cap int) StackOfPlates {
	return StackOfPlates{
		cap: cap,
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap <= 0 {
		return
	}

	last := this.GetLastStack()
	if last != nil && !last.isFull() {
		last.push(val)
	} else { // 初始时 last 为 nil，或者 last 已装满
		stack := NewSliceStack[int](this.cap)
		stack.push(val)
		this.stacks = append(this.stacks, stack)
	}
	fmt.Printf("stacks: %v\n", this.stacks)
}

func (this *StackOfPlates) Pop() int {
	if this.cap <= 0 {
		return -1
	}

	last := this.GetLastStack()
	if last == nil {
		return -1
	}
	top := last.pop()
	if last.isEmpty() {
		this.stacks = this.stacks[:len(this.stacks)-1]
	}
	return top
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.cap <= 0 || index < 0 || index >= len(this.stacks) {
		return -1
	}

	// 控制流不可能到达这里，因为 pop 过程中，一旦 stack 为空，就将其移除了
	// if this.stacks[index].isEmpty() {
	// 	return -1
	// }

	defer func() {
		fmt.Printf("stacks: %v\n", this.stacks)
	}()
	return this.leftShift(index, true)
}

// leetcode 期望的简单版本
func (this *StackOfPlates) PopAt2(index int) int {
	if this.cap <= 0 || index < 0 || index >= len(this.stacks) {
		return -1
	}

	s := this.stacks[index]
	r := s.pop()
	if s.isEmpty() {
		this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
	}
	return r
}

// 理解下面这个函数，需要根据题目要求推导出一个隐含的前提：
// 叠盘子时一个栈装满后再装下一个。pop 后若 last 栈为空，立即删除。
// 这就意味着，当从指定栈 pop 后，该栈只差一个就满，只需从下一个非
// last 栈补一个，下一个栈再从它的下一个补一个，以此类推，直到从
// last 栈补一个为止。如果 last 栈为空，就将其删除。
//
// 定义 removeTop 参数是为了将 popAt 函数中弹出元素的逻辑也放入
// 这个递归方法中，为什么要这么做呢？
// 为了实现从 last 开始，弹出栈底元素以补给前一个栈，每一次递归调用
// 就需要返回栈底元素。如果 pop_at 发起的第一次调用不返回元素，递归
// 函数的定义就无法统一
func (this *StackOfPlates) leftShift(index int, removeTop bool) int {
	stack := this.stacks[index]

	removed := 0
	if removeTop {
		// 从 popAt 方法发起的调用走入这个分支
		removed = stack.pop()
	} else {
		// 针对后续站的递归调用，走这个分支
		removed = stack.popBottom()
	}

	// 根据算法要求，可推知删除 stack 绝不可能发生在非 last stack 上。
	// 因为叠盘子时一个栈填满，再填下一个，不可能存在中间某个栈只有一个元素，
	// 且它下一个就是 last，last 中还有元素。因此，这里可以简化为直接删除
	// 最后一个栈。
	//
	// 为什么将下面两种情况组合为一个 if ... else if ... 块？
	// 因为对于非 last，走递归调用分支；对于 last，要么为空被删除，
	// 要么直接返回 removed，不可能走 elsif 分支
	if stack.isEmpty() {
		this.stacks = this.stacks[:len(this.stacks)-1]
	} else if index+1 < len(this.stacks) {
		v := this.leftShift(index+1, false)
		stack.push(v)
	}

	return removed
}

func (this *StackOfPlates) GetLastStack() *SliceStack[int] {
	if len(this.stacks) == 0 {
		return nil
	}
	return this.stacks[len(this.stacks)-1]
}

func TestStackOfPlates(t *testing.T) {
	s := NewStackOfPlates(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("=== pop ===\n")
	r1 := s.PopAt(0)
	fmt.Printf("r1: %v\n", r1)
	r2 := s.PopAt(0)
	fmt.Printf("r2: %v\n", r2)
}
