package path

import (
	"fmt"
	"strings"
)

type stack struct {
	buff []string
}

func (s *stack) path() string {
	return fmt.Sprintf("/%s", strings.Join(s.buff, "/"))
}

func (s *stack) push(c string) {
	s.buff = append(s.buff, c)
}

func (s *stack) pop() {
	s.buff = s.buff[:len(s.buff)-1]
}

func (s *stack) empty() bool {
	return len(s.buff) == 0
}

func split(s string) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			var v string
			for i := i + 1; i < len(s); i++ {
				if s[i] == '/' {
					break
				}
				v += string(s[i])
			}
			result = append(result, v)
		}
	}
	return result
}

func Simplify(s string) string {
	var st stack
	parts := split(s)
	for _, part := range parts {
		if part == "" {
			continue
		}
		if part == "." {
			continue
		}
		if part == ".." {
			if !st.empty() {
				st.pop()
			}
			continue
		}
		st.push(part)
	}
	return st.path()
}
