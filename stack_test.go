package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{}

	if s.Depth() != 0 {
		t.Fatalf("expected depth 0, got %d", s.Depth())
	}

	s.Push("one")
	s.Push("two")

	if s.Depth() != 2 {
		t.Fatalf("expected depth 2, got %d", s.Depth())
	}

	two := s.Pop()
	one := s.Pop()

	if s.Depth() != 0 {
		t.Fatalf("expected depth 0, got %d", s.Depth())
	}

	if two != "two" {
		t.Fatalf("expected 'two', got '%s'", two)
	}

	if one != "one" {
		t.Fatalf("expected 'one', got '%s'", one)
	}

}
