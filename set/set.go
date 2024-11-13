// Package set provides a basic set implementation
package set

import (
	"fmt"
	"sync"
)

// Set is a basic set implementation
type Set[T comparable] struct {
	elements map[T]struct{}
	mu       sync.RWMutex
}

// New creates a new set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]struct{}),
	}
}

// Add adds an element to the set
func (s *Set[T]) Add(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements[element] = struct{}{}
}

// Remove removes an element from the set
func (s *Set[T]) Remove(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.elements, element)
}

// Print prints out all the elements in the set
func (s *Set[T]) Print() {
	s.mu.RLock()
	for k := range s.elements {
		fmt.Printf("%v ", k)
	}
	s.mu.RUnlock()
	fmt.Println()
}

// Length returns the number of elements in the set
func (s *Set[T]) Length() int {
	return len(s.elements)
}

// Contains returns true if the set contains the element
func (s *Set[T]) Contains(elem T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.elements[elem]
	if !ok {
		return false
	}
	return true
}
