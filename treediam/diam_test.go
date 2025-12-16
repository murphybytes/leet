package treediam

import (
	"strconv"
	"testing"

	"github.com/murphybytes/leet/collections"
	"github.com/stretchr/testify/assert"
)

func TestTreeDiam(t *testing.T) {
	tcs := []struct {
		input  []any
		expect int
	}{
		{
			input:  []any{1, 2, 3, 4, 5},
			expect: 3,
		},
		{
			input:  []any{1, 2},
			expect: 1,
		},
		{
			input:  []any{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2},
			expect: 8,
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			root := collections.BuildTree(tc.input)
			assert.Equal(t, tc.expect, TreeDiam(root))
		})
	}
}
