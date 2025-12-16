package collections

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack[T]) Top() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}
