package stack

import (
	"fmt"
	"math"
)

// 这个算法的思路是给栈再添加一个成员 min_stack，顾名思义，这个栈用来保存较小的值。
// 具体工作原理如下：
// 1. 当栈为空时，入栈的数不管多大，它就是当前最小值，因此将其推入 min_stack。
// 2. 自此之后，入栈的数有三种情形：
// (a) 如果比当前最小值大，就不推入 min_stack。
// (b) 如果小于等于当前最小值，就推入 min_stack。
// 这样，MinStack 自底向上越来越小，就像一摞盘子自底向上，盘子越来越小。
// pop 时，如果栈顶值和 min stack 栈顶值不相等（这里绝不可能小于，只可能大于。
// 因为如果小于，那么 min stack 中必然存在与其相等、且在当前栈顶之上的值），
// 就不需要 min stack pop，只有相等时才需要执行 min stack pop

type MinStack struct {
	list     []int
	minStack []int
}

/** initialize your data structure here. */
// leetcode name: NewMinStack
func NewMinStack() *MinStack {
	return &MinStack{
		list:     make([]int, 0, 16),
		minStack: make([]int, 0, 8),
	}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if x <= this.GetMin() {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		panic(fmt.Errorf("empty stack"))
	}

	top := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]

	if top == this.GetMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

	//return top
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return math.MaxInt
	}
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
