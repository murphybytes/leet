package maxdiff

import (
	"strconv"
	"testing"

	"github.com/murphybytes/leet/collections"
	"github.com/stretchr/testify/assert"
)

func TestMaxAncestorDiff(t *testing.T) {
	tcs := []struct {
		input    []any
		expected int
	}{
		{
			input:    []any{8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13},
			expected: 7,
		},
		{
			input:    []any{1, nil, 2, nil, 0, 3},
			expected: 3,
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, MaxAncestorDiff(collections.BuildTree(tc.input)))
		})
	}
}
