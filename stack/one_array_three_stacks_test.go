package stack

import "fmt"

type TripleInOne struct {
	vals       []int // one array
	stackCount int
	stackCap   int
	stackSizes []int
}

// leetcode: Constructor
func NewTripleInOne(stackSize int) TripleInOne {
	return TripleInOne{
		vals:       make([]int, 3*stackSize),
		stackCount: 3,
		stackCap:   stackSize,
		stackSizes: []int{0, 0, 0},
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if stackNum < 0 || stackNum >= this.stackCount {
		panic(fmt.Errorf("illegal arg stackNum: %d", stackNum))
	}
	if this.isFull(stackNum) {
		// panic(fmt.Errorf("full stack-%d", stackNum))
		return
	}

	this.stackSizes[stackNum]++
	this.vals[this.indexOfTop(stackNum)] = value
}

func (this *TripleInOne) Pop(stackNum int) int {
	if stackNum < 0 || stackNum >= this.stackCount {
		panic(fmt.Errorf("illegal arg stackNum: %d", stackNum))
	}
	if this.IsEmpty(stackNum) {
		// panic(fmt.Errorf("empty stack-%d", stackNum))
		return -1
	}

	topIdx := this.indexOfTop(stackNum)
	v := this.vals[topIdx]
	this.vals[topIdx] = 0
	this.stackSizes[stackNum]--
	return v
}

func (this *TripleInOne) Peek(stackNum int) int {
	if stackNum < 0 || stackNum >= this.stackCount {
		panic(fmt.Errorf("illegal arg stackNum: %d", stackNum))
	}
	if this.IsEmpty(stackNum) {
		// panic(fmt.Errorf("empty stack-%d", stackNum))
		return -1
	}
	return this.vals[this.indexOfTop(stackNum)]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	if stackNum < 0 || stackNum >= this.stackCount {
		panic(fmt.Errorf("illegal arg stackNum: %d", stackNum))
	}
	return this.stackSizes[stackNum] == 0
}

func (this *TripleInOne) isFull(stackNum int) bool {
	return this.stackSizes[stackNum] == this.stackCap
}

func (this *TripleInOne) indexOfTop(stackNum int) int {
	offset := stackNum * this.stackCap
	size := this.stackSizes[stackNum]
	return offset + size - 1
}
