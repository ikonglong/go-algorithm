package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo2DArray(t *testing.T) {
	assert.Equal(t, [][]int{}, To2DArray("[]"))
	assert.Equal(t, [][]int{{0, 7}}, To2DArray("[[0,7]]"))
	assert.Equal(t, [][]int{{0, 7}, {0, 8}, {6, 1}}, To2DArray("[[0,7],[0,8],[6,1]]"))
}
