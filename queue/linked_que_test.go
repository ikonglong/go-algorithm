package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LinkedQueue[T any] struct {
	head *queNode[T]
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

type queNode[T any] struct {
	data T
	next *queNode[T]
}

func (q *LinkedQueue[T]) add(item T) {
	n := &queNode[T]{
		data: item,
	}

	if q.head == nil {
		q.head = n
		return
	}

	p := q.head
	for p.next != nil {
		p = p.next
	}
	p.next = n
}

func (q *LinkedQueue[T]) remove() T {
	if q.head == nil {
		panic(fmt.Errorf("empty queue"))
	}
	data := q.head.data
	q.head = q.head.next
	return data
}

func (q *LinkedQueue[T]) peek() T {
	if q.isEmpty() {
		panic(fmt.Errorf("empty queue"))
	}
	return q.head.data
}

func (q *LinkedQueue[T]) isEmpty() bool {
	return q.head == nil
}

func TestLinkedQueue(t *testing.T) {
	q := NewSliceQueue[int](1)
	assert.True(t, q.isEmpty())
	q.add(1)
	assert.False(t, q.isEmpty())
	q.add(2)
	q.add(3)
	assert.Equal(t, 1, q.remove())
	assert.Equal(t, 2, q.peek())
	q.remove()
	q.remove()
	// q.peek()
}
