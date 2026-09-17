package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LinkedStack[T any] struct {
	top *stackNode[T]
}

type stackNode[T any] struct {
	data T
	next *stackNode[T]
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) push(data T) {
	node := &stackNode[T]{
		data: data,
		next: s.top,
	}
	s.top = node
}

func (s *LinkedStack[T]) pop() T {
	if s.top == nil {
		panic(fmt.Errorf("empty stack"))
	}
	topData := s.top.data
	s.top = s.top.next
	return topData
}

func (s *LinkedStack[T]) peek() T {
	if s.top == nil {
		panic(fmt.Errorf("empty stack"))
	}
	return s.top.data
}

func (s *LinkedStack[T]) isEmpty() bool {
	return s.top == nil
}

func TestLinkedStack(t *testing.T) {
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
