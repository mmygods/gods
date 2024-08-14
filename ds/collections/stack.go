// Description: Stack interface.
package collections

// Stack represents a stack data structure.
type Stack[T any] interface {
	// Push adds an element to the top of the stack.
	Push(data T)
	// Pop removes and returns the element at the top of the stack.
	Pop() (T, bool)
	// Peek returns the element at the top of the stack without removing it.
	Peek() (T, bool)
	// IsEmpty returns true if the stack is empty, false otherwise.
	IsEmpty() bool
	// Length returns the number of elements in the stack.
	Length() int
}
