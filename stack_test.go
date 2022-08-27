package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Pop() != 3 {
		t.Fatalf(`pop should be 3`)
	}
	if s.Pop() != 2 {
		t.Fatalf(`pop should be 2`)
	}
	if s.Pop() != 1 {
		t.Fatalf(`pop should be 1`)
	}
}
