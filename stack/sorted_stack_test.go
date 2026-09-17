package stack

// 这个算法思路的心智模型是：
// 1. 想象有一个有序的 stack s，自底向上从大到小排列元素，或者说自顶向下从小到大排列元素
// 2. 从这个栈中间某个位置水平切分，把元素排列分成上下两部分，再把上半部分倒过来，装入一个辅助栈。
//    为什么要倒过来呢？因为对于栈，只能访问栈顶元素。插入新值到有序序列时，经常要把它插入在相
//    邻两个元素之间。把上半部分倒过来后，通过在两个栈之间移动元素，就可以将相邻元素分别放在两个
//    栈的栈顶。这样既方便了访问它们，又方便插入新值。
// 3. push 新值 val：通过在两个栈之间移动元素，使 val 小于等于 s 栈顶元素，大于等于辅助栈栈顶元素
//
// 跟下面的笨办法做下对比:
// 1. push val 时先将 s 中小于 val 的元素逐个弹出，再 push 进辅助栈
// 2. 将 value push 到 s
// 3. 先将辅助栈中的元素逐个弹出，再 push 进 s
//
// 这里采用的算法思路跟笨办法相比，将小于 val 的元素逐个弹出并 push 进辅助栈后，就放在那里，不会在
// 将 val 插入后马上就把它们再移回 s。这就是「延迟移动元素」策略，这个策略的应用挺广泛，例如著名的
// U&F算法、区间树、支持 min 操作的栈等就使用了它。

type SortedStack struct {
	s   []int
	aux []int
}

// leetcode name: Constructor
func NewSortedStack() SortedStack {
	return SortedStack{
		s:   make([]int, 0),
		aux: make([]int, 0),
	}
}

func (this *SortedStack) Push(val int) {
	for len(this.s) != 0 && val > this.s[len(this.s)-1] {
		this.aux = append(this.aux, this.s[len(this.s)-1])
		this.s = this.s[:len(this.s)-1]
	}
	for len(this.aux) != 0 && val < this.aux[len(this.aux)-1] {
		this.s = append(this.s, this.aux[len(this.aux)-1])
		this.aux = this.aux[:len(this.aux)-1]
	}

	// 如果 s 和 aux 都为空，或者通过移动元素已将插入位置下上两个元素分别放在了两个栈的栈顶
	this.s = append(this.s, val) // ok: this.aux = append(this.aux, val)
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		// panic(fmt.Errorf("empty stack"))
		return
	}
	this.shiftAuxToS()
	this.s = this.s[:len(this.s)-1]
}

func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		// panic(fmt.Errorf("empty stack"))
		return -1
	}
	this.shiftAuxToS()
	return this.s[len(this.s)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.s)+len(this.aux) == 0
}

func (this *SortedStack) shiftAuxToS() {
	for len(this.aux) != 0 {
		// pop from aux, push to s
		top := this.aux[len(this.aux)-1]
		this.s = append(this.s, top)
		this.aux = this.aux[:len(this.aux)-1]
	}
}
