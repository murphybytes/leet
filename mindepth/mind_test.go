package mindepth

import (
	"strconv"
	"testing"

	"github.com/murphybytes/leet/collections"
	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tcs := []struct {
		input    []any
		expected []int
	}{
		{
			input:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: []int{9, 15, 7, 20, 3},
		},
		{
			input:    []any{2, nil, 3, nil, 4, nil, 5, nil, 6},
			expected: []int{6, 5, 4, 3, 2},
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			root := collections.BuildTree(tc.input)
			actual := postorder(root)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func postorder(root *collections.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	left := postorder(root.Left)
	if left != nil {
		result = append(result, left...)
	}
	right := postorder(root.Right)
	if right != nil {
		result = append(result, right...)
	}
	result = append(result, root.Val)

	return result
}
