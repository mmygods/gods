// Description: This package contains the implementation of stack data structure.
package stack

import (
	"github.com/mmygods/gods/ds/collections"
	"github.com/mmygods/gods/ds/models/dll"
)

type Stack[T any] struct {
	data collections.List[T]
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

// New creates a new stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{data: &dll.DoublyLinkedList[T]{}}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(data T) {
	s.data.Append(data)
}

// Pop removes and returns the element at the top of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if s.data.IsEmpty() {
		return zeroValue[T](), false
	}
	return s.data.Pop()
}

// Peek returns the element at the top of the stack without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if s.data.IsEmpty() {
		return zeroValue[T](), false
	}
	return s.data.Get(s.data.Length() - 1)
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.data.IsEmpty()
}

// Length returns the number of elements in the stack.
func (s *Stack[T]) Length() int {
	return s.data.Length()
}
