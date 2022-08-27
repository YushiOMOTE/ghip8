package main

type Stack struct {
	sp    int
	stack []uint16
}

func NewStack() Stack {
	return Stack{
		sp:    0,
		stack: make([]uint16, 32),
	}
}

func (s *Stack) Push(v uint16) {
	s.stack[s.sp] = v
	s.sp += 1
}

func (s *Stack) Pop() uint16 {
	s.sp -= 1
	return s.stack[s.sp]
}
