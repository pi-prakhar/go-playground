package helpers

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	topIndex := len(s.elements) - 1
	element := s.elements[topIndex]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) IsEmpty() (bool) {
	return len(s.elements) == 0 
}
