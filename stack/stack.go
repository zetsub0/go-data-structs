// Package stack provides a basic stack implementation
package stack

import (
	"fmt"
)

// Stack is a basic stack implementation
type Stack[T comparable] struct {
	values []T
}

// Push adds a value to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// Pop removes the top value from the stack and returns it.
// If the stack is empty, it returns the zero value of type T
// and an error.
func (s *Stack[T]) Pop() (T, error) {
	const fn = "Pop"
	if s.IsEmpty() {
		var nilVal T
		err := fmt.Errorf("%s: stack is empty", fn)

		return nilVal, err
	}
	popVal := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return popVal, nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// GetSize returns the size of the stack
func (s *Stack[T]) GetSize() int {
	return len(s.values)
}
