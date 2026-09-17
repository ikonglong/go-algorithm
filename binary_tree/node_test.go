package binary_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	levelOrderList := []any{3, 9, 20, "null", "null", 15, 7}
	root := NewTree(levelOrderList)
	vList := ValList(root)
	fmt.Printf("vList:%v\n", vList)
	assert.Equal(t, levelOrderList, vList)
}

func TestValList(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	vList := ValList(root)
	assert.Equal(t,
		[]any{3, 9, 20, "null", "null", 15, 7},
		vList)

	vList2 := ValList2(root)
	assert.Equal(t,
		[]any{3, 9, 20, "null", "null", 15, 7},
		vList2)

	// vList22 := ValList2(root)
	// assert.Equal(t,
	// 	[]any{3, 9, 20, "null", "null", 15, 7, "null", "null", "null", "null"},
	// 	vList22)
}
