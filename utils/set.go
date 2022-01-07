package utils

var exists = struct{}{}

type Set struct {
	values map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		values: make(map[int]struct{}),
	}
}

func (s *Set) Add(value int) {
	s.values[value] = exists
}

func (s *Set) Remove(value int) {
	delete(s.values, value)
}

func (s *Set) Contains(value int) bool {
	_, ok := s.values[value]
	return ok
}

func (s *Set) Values() []int {
	keys := make([]int, 0, len(s.values))
	for k := range s.values {
		keys = append(keys, k)
	}

	return keys
}
