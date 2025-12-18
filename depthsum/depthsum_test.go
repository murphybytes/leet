package depthsum

import (
	"strconv"
	"testing"

	"github.com/murphybytes/leet/collections"
	"github.com/stretchr/testify/assert"
)

func TestDeepestLeavesSum(t *testing.T) {
	tcs := []struct {
		input    []any
		expected int
	}{
		{
			input:    []any{1, 2, 3, 4, 5, nil, 6, 7, nil, nil, nil, nil, 8},
			expected: 15,
		},
		{
			input:    []any{6, 7, 8, 2, 7, 1, 3, 9, nil, 1, 4, nil, nil, nil, 5},
			expected: 19,
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			root := collections.BuildTree(tc.input)
			actual := DeepestLeavesSum(root)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
