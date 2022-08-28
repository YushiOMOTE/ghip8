package main

// Represents a stack with 16bits items
type Stack struct {
	sp    int
	stack []uint16
}

// Creates a new Stack instance
func NewStack() Stack {
	return Stack{
		sp:    0,
		stack: make([]uint16, 32),
	}
}

// Pushes an item to the stack
func (s *Stack) Push(v uint16) {
	s.stack[s.sp] = v
	s.sp += 1
}

// Pops an item from the stack
func (s *Stack) Pop() uint16 {
	s.sp -= 1
	return s.stack[s.sp]
}
