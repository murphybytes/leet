package nge

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	tcs := []struct {
		ar1      []int
		ar2      []int
		expected []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, NextGreaterElement(tc.ar1, tc.ar2), tc.expected)
		})
	}
}
