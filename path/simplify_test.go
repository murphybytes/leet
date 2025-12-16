package path

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplify(t *testing.T) {
	tcs := []struct {
		input, expected string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
	}

	for i, tc := range tcs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, Simplify(tc.input))
		})
	}
}
