package main

type Stack struct {
	Items []Color
}

func (s *Stack) Push(item Color) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() Color {
	if len(s.Items) == 0 {
		panic("underflow")
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func (s *Stack) Depth() int {
	return len(s.Items)
}

func (s *Stack) Top() Color {
	if len(s.Items) == 0 {
		panic("underflow")
	}
	item := s.Items[len(s.Items)-1]
	return item
}
