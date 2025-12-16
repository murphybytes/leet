package gs

import (
	"github.com/murphybytes/leet/collections"
)

func MakeGood(s string) string {
	var stack collections.ByteStack

	buff := []byte(s)

	for _, v := range buff {
		if stack.Empty() {
			stack.Push(v)
			continue
		}

		if same(stack.Top(), v) {
			stack.Pop()
			continue
		}
		stack.Push(v)
	}

	return stack.String()
}

func same(l, r byte) bool {
	i := int(l)
	j := int(r)
	k := i - j
	return k == 32 || k == -32
}
