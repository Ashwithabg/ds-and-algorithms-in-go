package main

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) push(value int) {
	s.data = append(s.data, value)
}

func (s *stack) pop() int {

	lastElement := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return lastElement
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}
