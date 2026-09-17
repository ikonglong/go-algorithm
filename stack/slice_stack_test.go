package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SliceStack[T any] struct {
	list []T
	cap  int
}

// 虽然定义了容量（capacity）参数，但目前并未实现，容量限制可以无限扩展。
// 提供了 isFull 方法，仅用来查看是否到达指定容量了
func NewSliceStack[T any](cap int) *SliceStack[T] {
	return &SliceStack[T]{
		list: make([]T, 0, cap),
		cap:  cap,
	}
}

func (s *SliceStack[T]) push(item T) {
	s.list = append(s.list, item)
}

func (s *SliceStack[T]) pop() T {
	if s.isEmpty() {
		panic(fmt.Errorf("empty stack"))
	}
	top := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return top
}

func (s *SliceStack[T]) peek() T {
	if s.isEmpty() {
		panic(fmt.Errorf("empty stack"))
	}
	return s.list[len(s.list)-1]
}

func (s *SliceStack[T]) popBottom() T {
	if s.isEmpty() {
		panic(fmt.Errorf("empty stack"))
	}
	bottom := s.list[0]
	s.list = s.list[1:]
	return bottom
}

func (s *SliceStack[T]) isEmpty() bool {
	return len(s.list) == 0
}

func (s *SliceStack[T]) isFull() bool {
	return len(s.list) == s.cap
}

func (s *SliceStack[T]) len() int {
	return len(s.list)
}

func (s *SliceStack[T]) String() string {
	return fmt.Sprintf("%v", s.list)
}

func TestSliceStack(t *testing.T) {
	s := NewSliceStack[int](3)
	assert.True(t, s.isEmpty())
	s.push(1)
	assert.False(t, s.isEmpty())
	assert.Equal(t, 1, s.peek())
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	assert.Equal(t, 5, s.peek())
	for i := 1; i <= 5; i++ {
		v := s.pop()
		assert.Equal(t, 6-i, v)
	}
	// s.pop()
	s.peek()
}
