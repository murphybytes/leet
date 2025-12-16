package collections

type ByteStack struct {
	buff []byte
}

func (s *ByteStack) Push(c byte) {
	s.buff = append(s.buff, c)
}

func (s *ByteStack) Pop() {
	s.buff = s.buff[:len(s.buff)-1]
}

func (s *ByteStack) Empty() bool {
	return len(s.buff) == 0
}

func (s *ByteStack) Top() byte {
	return s.buff[len(s.buff)-1]
}

func (s *ByteStack) String() string {
	return string(s.buff)
}
