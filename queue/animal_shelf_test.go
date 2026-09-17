package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	CAT = iota
	DOG
)

type Animal struct {
	no          int
	shelteredAt int // 模拟收容时间
}

func (a *Animal) isOlderThan(b *Animal) bool {
	return a.shelteredAt < b.shelteredAt
}

type AnimalShelf struct {
	dogQue      []*Animal
	catQue      []*Animal
	shelteredAt int
}

// leetcode name: Constructor
func NewAnimalShelf() AnimalShelf {
	return AnimalShelf{
		dogQue: make([]*Animal, 0),
		catQue: make([]*Animal, 0),
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	// 从 1 开始或者创建时将 shelter 初始化为 -1。
	// 这行代码跟使用它执行结果的代码写在一起，就不容易埋 bug
	this.shelteredAt++
	a := &Animal{
		no:          animal[0],
		shelteredAt: this.shelteredAt,
	}
	if animal[1] == CAT { // cat
		this.catQue = append(this.catQue, a)
	} else if animal[1] == DOG { // dog
		this.dogQue = append(this.dogQue, a)
	}
	// 从零开始，入队后，自增加这行代码就很容易漏写，容易埋 bug
	// this.shelteredAt++
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.catQue) == 0 && len(this.dogQue) == 0 {
		return []int{-1, -1}
	} else if (len(this.catQue)) == 0 {
		return this.DequeueDog()
	} else if len(this.dogQue) == 0 {
		return this.DequeueCat()
	}

	// cat 和 dog 队列都不为空
	// 类似将两个有序队列归并为一个有序队列
	if this.catQue[0].isOlderThan(this.dogQue[0]) {
		return this.DequeueCat()
	} else {
		return this.DequeueDog()
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dogQue) == 0 {
		return []int{-1, -1}
	}
	head := this.dogQue[0]
	this.dogQue = this.dogQue[1:]
	return []int{head.no, DOG}
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.catQue) == 0 {
		return []int{-1, -1}
	}
	head := this.catQue[0]
	this.catQue = this.catQue[1:]
	return []int{head.no, CAT}
}

func TestAnimalShelf(t *testing.T) {
	s := NewAnimalShelf()
	s.Enqueue([]int{0, 0})
	s.Enqueue([]int{1, 0})
	assert.Equal(t, []int{0, 0}, s.DequeueCat())
	assert.Equal(t, []int{-1, -1}, s.DequeueDog())
	assert.Equal(t, []int{1, 0}, s.DequeueAny())

	s2 := NewAnimalShelf()
	s2.Enqueue([]int{0, 0})
	s2.Enqueue([]int{1, 0})
	s2.Enqueue([]int{2, 1})
	assert.Equal(t, []int{2, 1}, s2.DequeueDog())
	assert.Equal(t, []int{0, 0}, s2.DequeueCat())
	assert.Equal(t, []int{1, 0}, s2.DequeueAny())

	s3 := NewAnimalShelf()
	s3.Enqueue([]int{0, 0})
	s3.Enqueue([]int{1, 1})
	s3.Enqueue([]int{2, 0})
}
