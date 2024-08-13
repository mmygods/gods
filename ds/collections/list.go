package collections

import "github.com/mmygods/gods/ds/models/dll"

type Node[T any] interface {
	GetData() T
}

// List is a collection of elements that can be accessed by index.
type List[T any] interface {
	// Append adds an element to the end of the list.
	Append(T) bool
	// Prepend adds an element to the beginning of the list.
	Prepend(T) bool
	// Insert adds an element at the specified index.
	Insert(int, T) bool
	// Get returns the element at the specified index.
	Get(int) (T, bool)
	// Set sets the element at the specified index.
	Set(int, T) bool
	// Delete removes the element at the specified index.
	Delete(int) bool
	// Length returns the number of elements in the list.
	Length() int
	// IsEmpty returns true if the list is empty.
	IsEmpty() bool
	// Pop removes and returns the last element in the list.
	Pop() (T, bool)
	// PopFirst removes and returns the first element in the list.
	PopFirst() (T, bool)
	// Range returns a channel that iterates over the elements in the list.
	Range() <-chan T
}

type LruList[T any, N Node[T]] interface {
	List[T]
	// AppendNode adds a node to the end of the list.
	AppendNode(*dll.DllNode[T]) bool
	// PrependNode adds a node to the beginning of the list.
	PrependNode(*dll.DllNode[T]) bool
	// PopNode removes and returns the last node in the list.
	PopNode() (*dll.DllNode[T], bool)
	// PopFirstNode removes and returns the first node in the list.
	PopFirstNode() (*dll.DllNode[T], bool)
	// GetNode returns the node at the specified index.
	GetNode(int) *dll.DllNode[T]
	// DeleteNode removes the node at the specified index.
	DeleteNode(*dll.DllNode[T]) bool
}
