package queue

import "testing"

// 用两个栈 stackNewest, stackOldest 来实现队列。
// 简单版本：
//
//	入队时，将新元素添加到 stack_newest 栈顶
//	出队时，将 stack_newest 倒入 stack_oldest，移除栈顶元素，再从 stack oldest 倒回 stack newest
//	peek 时，将 stack newest 倒入 stack oldest，返回栈顶元素，再从 stack oldest 倒回 stack newest
//
// 改进版本：
// 如果连续执行做出队/peek 操作，就会在两个栈之间来回倒元素。
// 若从 newest 倒入 oldest 后，就保留在那里，不再倒回。让 newest 保存新近添加的元素，让 oldest 保存
// 久远添加的元素，就不用来回倒元素了。
//
// 这个思路挺巧妙，它的心智模型是，逻辑上将 newest 的栈底接上 oldest 的栈底，就把两个站合并成了一个队列。
// 此外，这里还用到了「延迟移动元素」策略来改进性能。
//

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
type MyQueue struct {
	stackNewest []int
	stackOldest []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackNewest: make([]int, 0),
		stackOldest: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackNewest = append(this.stackNewest, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.shiftStack()
	// this.shiftStack2()
	head := this.stackOldest[len(this.stackOldest)-1]
	this.stackOldest = this.stackOldest[:len(this.stackOldest)-1]
	// 如果是非基础类型的值，不置为 nil，可能导致内存泄漏。
	// 但是，非基础类型的数组或切片，也不能将对应位置设置为 0，只能设置为零值。
	// this.stackOldest[len(this.stackOldest)-1] = 0
	// this.shiftStack2()
	return head
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.shiftStack()
	// this.shiftStack2()
	head := this.stackOldest[len(this.stackOldest)-1]
	// this.shiftStack2()
	return head
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stackNewest)+len(this.stackOldest) == 0
}

func (this *MyQueue) shiftStack() {
	if len(this.stackOldest) == 0 {
		// 原本是为了防止底层数组内存泄漏，所以才新创建一个切片，
		// 但这里是基础类型，所以不存在内存泄漏
		// this.stackOldest = make([]int, 0)
		for i := len(this.stackNewest) - 1; i >= 0; i-- {
			this.stackOldest = append(this.stackOldest, this.stackNewest[i])
		}
		this.stackNewest = this.stackNewest[:0]
	}
}

func (this *MyQueue) shiftStack2() {
	if len(this.stackOldest) == 0 {
		// 原本是为了防止底层数组内存泄漏，所以才新创建一个切片，
		// 但这里是基础类型，所以不存在内存泄漏
		// this.stackOldest = make([]int, 0)
		for i := len(this.stackNewest) - 1; i >= 0; i-- {
			this.stackOldest = append(this.stackOldest, this.stackNewest[i])
		}
		this.stackNewest = this.stackNewest[:0]
	} else if len(this.stackNewest) == 0 {
		// 原本是为了防止底层数组内存泄漏，所以才新创建一个切片，
		// 但这里是基础类型，所以不存在内存泄漏
		// this.stackNewest = make([]int, 0)
		for i := len(this.stackOldest) - 1; i >= 0; i-- {
			this.stackNewest = append(this.stackNewest, this.stackOldest[i])
		}
		this.stackOldest = this.stackOldest[:0]
	}
}

func TestMyQueue(t *testing.T) {}
