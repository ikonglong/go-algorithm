package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 排序给定栈的算法思路如下：
//  1. 引入一个辅助栈，假定辅助栈总是自底向上从大到小有序的。初始时是空的，因此也是有序的
//  2. 只要 S 不为空，
//     就从 S 弹出一个元素，存入 tmp
//     只要 tmp 比辅助栈栈顶元素大，
//     就将辅助栈中的元素弹出并放入 S，
//     将新元素压入辅助栈
//  3. 将辅助栈中的元素倒入 S
func sort(s *SliceStack[int]) {
	if s.len() == 0 || s.len() == 1 {
		return
	}

	aux := NewSliceStack[int](16)
	for !s.isEmpty() {
		tmp := s.pop()
		// 为保持辅助栈有序而暂时将大于 tmp 的元素转移到 S，
		// 插入 tmp 后，会继续遍历 S，这些元素又自然转移回来了，
		// 并且顺序刚好是期望的
		for !aux.isEmpty() && tmp > aux.peek() {
			s.push(aux.pop())
		}
		aux.push(tmp)
	}

	for !aux.isEmpty() {
		s.push(aux.pop())
	}
}

func TestSortStack(t *testing.T) {
	s := NewSliceStack[int](8)
	sort(s)
	assert.Equal(t, []int{}, s.list)

	s2 := NewSliceStack[int](8)
	s2.push(2)
	s2.push(1)
	sort(s2)
	assert.Equal(t, []int{1, 2}, s2.list)

	s3 := NewSliceStack[int](8)
	s3.push(5)
	s3.push(6)
	s3.push(3)
	s3.push(4)
	s3.push(1)
	sort(s3)
	assert.Equal(t, []int{1, 3, 4, 5, 6}, s3.list)

	s4 := NewSliceStack[int](8)
	s4.push(3)
	s4.push(2)
	s4.push(1)
	sort(s4)
	assert.Equal(t, []int{1, 2, 3}, s4.list)

	s5 := NewSliceStack[int](8)
	s5.push(1)
	s5.push(2)
	s5.push(3)
	sort(s5)
	assert.Equal(t, []int{1, 2, 3}, s5.list)
}
