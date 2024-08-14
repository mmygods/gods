package collections

// Deque represents a deque data structure.
type Deque[T any] interface {
	// Append adds an element to the end of the deque.
	Append(data T) bool
	// Prepend adds an element to the beginning of the deque.
	Prepend(data T) bool
	// Pop removes and returns the element at the end of the deque.
	Pop() (T, bool)
	// PopFirst removes and returns the element at the beginning of the deque.
	PopFirst() (T, bool)
	// IsEmpty returns true if the deque is empty, false otherwise.
	IsEmpty() bool
	// Length returns the number of elements in the deque.
	Length() int
}
