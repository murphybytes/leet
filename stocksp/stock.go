package stocksp

import "math"

type stack []int

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

type StockSpanner struct {
	days  int
	stack stack
	vals  map[int]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []int{0},
		vals: map[int]int{
			0: math.MaxInt,
		},
	}
}

func (ss *StockSpanner) Next(price int) int {
	for price >= ss.vals[ss.stack.top()] {
		idx := ss.stack.pop()
		delete(ss.vals, idx)
	}

	ss.days += 1
	begin := ss.stack.top()
	ss.vals[ss.days] = price
	ss.stack.push(ss.days)

	return ss.days - begin
}
