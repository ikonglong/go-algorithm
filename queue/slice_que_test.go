package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SliceQueue[T any] struct {
	list []T
}

func NewSliceQueue[T any](cap int) *SliceQueue[T] {
	return &SliceQueue[T]{
		list: make([]T, 0, cap),
	}
}

func (q *SliceQueue[T]) add(item T) {
	q.list = append(q.list, item)
}

func (q *SliceQueue[T]) remove() T {
	if q.isEmpty() {
		panic(fmt.Errorf("empty queue"))
	}
	head := q.list[0]
	q.list = q.list[1:]
	return head
}

func (q *SliceQueue[T]) peek() T {
	if q.isEmpty() {
		panic(fmt.Errorf("empty queue"))
	}
	return q.list[0]
}

func (q *SliceQueue[T]) isEmpty() bool {
	return len(q.list) == 0
}

func TestSliceQueue(t *testing.T) {
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
	q.peek()
}
