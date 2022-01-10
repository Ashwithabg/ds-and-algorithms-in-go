package main

type stack struct {
	value []int
}

func newStack() *stack {
	return &stack{
		value: make([]int, 0),
	}
}

func (s *stack) push(value int) {
	s.value = append(s.value, value)
}

func (s *stack) pop() int {
	lastElement := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return lastElement
}

func (s stack) isEmpty() bool {
	return len(s.value) == 0
}
