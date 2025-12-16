package gs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGood(t *testing.T) {
	tcs := []struct {
		input, expected string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
		{"s", "s"},
		{"Pp", ""},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, MakeGood(tc.input))
		})
	}
}
