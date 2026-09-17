package linkedlist

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	linked := NewLinked(list)
	vList := linked.ValList()
	assert.Equal(t, list, vList)
}

func TestToLinkedList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	linkedlist := NewLinked(list)
	n := linkedlist
	list2 := make([]int, 0, 5)
	for n != nil {
		list2 = append(list2, n.Val)
		n = n.Next
	}
	assert.Equal(t, list, list2)
	assert.True(t, slices.Equal(list, list2))
}
