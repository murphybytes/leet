package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tcs := []struct {
		input    []any
		expected []int
	}{
		{
			input:    []any{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13},
			expected: []int{1, 3, 4, 6, 7, 8, 10, 13, 14},
		},
		{
			input:    []any{1, nil, 2, nil, 0, 3},
			expected: []int{1, 2, 3, 0},
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			root := BuildTree(tc.input)
			actual := df(root)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func df(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, df(root.Left)...)
	result = append(result, root.Val)
	result = append(result, df(root.Right)...)
	return result
}
