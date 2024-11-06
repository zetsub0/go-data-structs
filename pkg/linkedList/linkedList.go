package linkedList

import (
	"fmt"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

type LinkedList[T comparable] struct {
	head   *node[T]
	length int
}

// New returns a new empty LinkedList.
func New[T comparable](values []T) *LinkedList[T] {
	ll := &LinkedList[T]{}
	for _, value := range values {
		ll.Append(value)
	}
	return ll
}

// Append adds an item to the end of the list.
func (l *LinkedList[T]) Append(v T) {
	newNode := &node[T]{value: v}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.length++
}

// Prepend adds an item to the start of the list.
func (l *LinkedList[T]) Prepend(v T) {
	newNode := &node[T]{value: v, next: l.head}
	l.head = newNode
	l.length++
}

// Insert adds an item to the list at the given position.
// It returns an error if the position is out of range.
func (l *LinkedList[T]) Insert(v T, position int) error {
	const fn = "Insert"

	if position < 0 || position > l.length {
		return fmt.Errorf("%s: position %d out of range", fn, position)
	}

	if position == 0 {
		l.Prepend(v)
		return nil
	}

	if position == l.length {
		l.Append(v)
		return nil
	}

	prevNode, _ := l.getNodeAt(position - 1)
	newNode := &node[T]{value: v, next: prevNode.next}
	prevNode.next = newNode
	l.length++

	return nil

}

// RemoveHead removes head node from Linked list.
func (l *LinkedList[T]) RemoveHead() error {
	const fn = "RemoveHead"

	if l.length == 0 {
		return fmt.Errorf("%s: list is empty", fn)
	}

	current := l.head.next
	l.head = current
	l.length--
	return nil
}

// RemoveTail removes the tail node of the list.
func (l *LinkedList[T]) RemoveTail() error {
	const fn = "RemoveTail"

	if l.length == 0 {
		return fmt.Errorf("%s: list is empty", fn)
	}
	if l.length == 1 {
		return l.RemoveHead()
	}

	secondLast, err := l.getNodeAt(l.length - 2)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	secondLast.next = nil
	l.length--

	return nil
}

// RemoveNode removes specified node from Linked list.
func (l *LinkedList[T]) RemoveNode(position int) error {
	if position < 0 || position >= l.length {
		return fmt.Errorf("RemoveNode: position %d out of range", position)
	}
	if position == 0 {
		return l.RemoveHead()
	}
	if position == l.length-1 {
		return l.RemoveTail()
	}

	prevNode, err := l.getNodeAt(position - 1)
	if err != nil {
		return fmt.Errorf("RemoveNode: %w", err)
	}
	prevNode.next = prevNode.next.next
	l.length--

	return nil
}

// Print prints out all the values in the list. This is mostly useful for
// debugging.
func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// SearchItem searches for an item in the list.
func (l *LinkedList[T]) SearchItem(value T) bool {
	current := l.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

// GetLength prints the length of the list.
func (l *LinkedList[T]) GetLength() int {
	return l.length
}

func (l *LinkedList[T]) getNodeAt(position int) (*node[T], error) {
	const fn = "getNodeAt"
	if position == 0 || position > l.length {
		return nil, fmt.Errorf("%s: position %d out of range", fn, position)
	}

	current := l.head

	for i := 1; i < position; i++ {
		current = current.next
	}

	return current, nil
}
