package stocksp

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockSpanner(t *testing.T) {
	tcs := []struct {
		prices   []int
		expected []int
	}{
		{
			[]int{100, 80, 60, 70, 60, 75, 85},
			[]int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			[]int{28, 13, 28, 35, 48, 53, 66, 80, 87, 88},
			[]int{1, 1, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ss := Constructor()
			for j, price := range tc.prices {
				span := ss.Next(price)
				assert.Equal(t, span, tc.expected[j])
			}
		})
	}
}
