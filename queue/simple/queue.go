// Package queue provides a simple queue implementation
package queue

import (
	"errors"
	"fmt"
)

// Queue is a simple queue implementation
type Queue[T any] struct {
	maxLength int
	values    []T
	length    int
}

// New creates a new queue
func New[T any](length int) *Queue[T] {
	return &Queue[T]{
		maxLength: length,
		values:    make([]T, 0, length),
	}
}

// Enqueue adds a value to the end of the queue
func (q *Queue[T]) Enqueue(value T) error {
	const fn = "Enqueue"

	if q.length < q.maxLength {
		q.values = q.values[:q.length+1]
		q.values[q.length] = value
		q.length++
	} else {
		err := errors.New(errFullQueue)
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

// Rear returns the element at the rear end without removing it
func (q *Queue[T]) Rear() (T, error) {
	const fn = "Rear"
	if q.IsEmpty() {
		var nilVal T
		return nilVal, fmt.Errorf("%s: %v", fn, errEmptyQueue)
	}

	return q.values[q.length-1], nil
}

// Dequeue removes the first value from the queue and returns it
func (q *Queue[T]) Dequeue() (T, error) {
	const fn = "Dequeue"
	if q.length == 0 {
		var nilVal T
		return nilVal, fmt.Errorf("%s: %v", fn, errEmptyQueue)
	} else {
		val := q.values[0]
		q.values = q.values[1:]
		q.length--
		return val, nil
	}
}

// Front returns the first value in the queue without removing it
func (q *Queue[T]) Front() (T, error) {
	const fn = "Front"
	if q.length == 0 {
		var nilVal T
		return nilVal, fmt.Errorf("%s: %v", fn, errEmptyQueue)
	} else {
		return q.values[0], nil
	}
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

// Print prints all values in the queue
func (q *Queue[T]) Print() {
	for _, v := range q.values {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// GetLength returns the length of the queue
func (q *Queue[T]) GetLength() int {
	return q.length
}
