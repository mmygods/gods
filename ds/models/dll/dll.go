// Description: This package contains the implementation of a doubly linked list.

package dll

import "sync"

// The DllNode struct represents a node in a doubly linked list.
type DllNode[T any] struct {
	data T
	next *DllNode[T]
	prev *DllNode[T]
}

// The DoublyLinkedList struct represents a doubly linked list.
type DoublyLinkedList[T any] struct {
	length int
	head   *DllNode[T]
	tail   *DllNode[T]
	mu     sync.RWMutex
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

// NewNode creates a new node with the specified data.
func NewNode[T any](data T) *DllNode[T] {
	return &DllNode[T]{data: data}
}

// GetData returns the data stored in the node.
func (node DllNode[T]) GetData() T {
	return node.data
}

// append adds an element to the end of the list.
func (dll *DoublyLinkedList[T]) append(data T) bool {
	node := &DllNode[T]{data: data}
	return dll.appendNode(node)
}

// appendNode adds a node to the end of the list.
func (dll *DoublyLinkedList[T]) appendNode(node *DllNode[T]) bool {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.prev = dll.tail
		dll.tail.next = node
		dll.tail = node
	}
	dll.length++
	return true
}

// prepend adds a node to the beginning of the list.
func (dll *DoublyLinkedList[T]) prepend(data T) bool {
	node := &DllNode[T]{data: data}
	return dll.prependNode(node)
}

// prependNode adds a node to the beginning of the list.
func (dll *DoublyLinkedList[T]) prependNode(node *DllNode[T]) bool {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
	dll.length++
	return true
}

// pop removes and returns the last element in the list.
func (dll *DoublyLinkedList[T]) pop() (T, bool) {
	if node, ok := dll.popNode(); ok {
		return node.data, true
	} else {
		return zeroValue[T](), false
	}
}

// popNode removes and returns the last node in the list.
func (dll *DoublyLinkedList[T]) popNode() (*DllNode[T], bool) {
	if dll.head == nil {
		return nil, false
	}
	node := dll.tail
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.length--
	node.next = nil
	node.prev = nil
	return node, true
}

// popFirst removes and returns the first element in the list.
func (dll *DoublyLinkedList[T]) popFirst() (T, bool) {
	if node, ok := dll.popFirstNode(); ok {
		return node.data, true
	} else {
		return zeroValue[T](), false
	}

}

// popFirstNode removes and returns the first element in the list.
func (dll *DoublyLinkedList[T]) popFirstNode() (*DllNode[T], bool) {
	if dll.head == nil {
		return nil, false
	}
	node := dll.head
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.length--
	node.next = nil
	node.prev = nil
	return node, true
}

// getNode returns the node at the specified index.
func (dll *DoublyLinkedList[T]) getNode(index int) *DllNode[T] {
	if index < 0 || index >= dll.length {
		return nil
	}
	node := dll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// get returns the element at the specified index.
func (dll *DoublyLinkedList[T]) get(index int) (T, bool) {
	node := dll.getNode(index)
	if node == nil {
		return zeroValue[T](), false
	}
	return node.data, true
}

// set sets the element at the specified index.
func (dll *DoublyLinkedList[T]) set(index int, data T) bool {
	node := dll.getNode(index)
	if node == nil {
		return false
	}
	node.data = data
	return true
}

// insert adds an element at the specified index.
func (dll *DoublyLinkedList[T]) insert(index int, data T) bool {
	if index < 0 || index > dll.length {
		return false
	}
	if index == 0 {
		return dll.prepend(data)
	}
	if index == dll.length {
		return dll.append(data)
	}
	node := &DllNode[T]{data: data}
	prevNode := dll.getNode(index - 1)
	nextNode := prevNode.next
	prevNode.next = node
	node.prev = prevNode
	node.next = nextNode
	nextNode.prev = node
	dll.length++
	return true
}

// delete removes the element at the specified index.
func (dll *DoublyLinkedList[T]) delete(index int) bool {
	node := dll.getNode(index)
	if node == nil {
		return false
	}
	return dll.deleteNode(node)
}

func (dll *DoublyLinkedList[T]) deleteNode(node *DllNode[T]) bool {
	if node == nil {
		return false
	}
	if node != dll.head && node != dll.tail && node.prev == nil && node.next == nil {
		return false
	}
	if node == dll.head {
		_, ok := dll.popFirst()
		return ok
	}
	if node == dll.tail {
		_, ok := dll.pop()
		return ok
	}
	prevNode := node.prev
	nextNode := node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
	node.next = nil
	node.prev = nil
	dll.length--
	return true
}

func (dll *DoublyLinkedList[T]) isEmpty() bool {
	return dll.length == 0
}

func (dll *DoublyLinkedList[T]) length_internal() int {
	return dll.length
}

// lock locks the mutex for writing.
func (dll *DoublyLinkedList[T]) lock() {
	dll.mu.Lock()
}

// unlock unlocks the mutex for writing.
func (dll *DoublyLinkedList[T]) unlock() {
	dll.mu.Unlock()
}

// rLock locks the mutex for reading.
func (dll *DoublyLinkedList[T]) rLock() {
	dll.mu.RLock()
}

// rUnlock unlocks the mutex for reading.
func (dll *DoublyLinkedList[T]) rUnlock() {
	dll.mu.RUnlock()
}

// Append adds an element to the end of the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Append(data T) bool {
	dll.lock()
	defer dll.unlock()
	return dll.append(data)
}

// Prepend adds a node to the beginning of the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Prepend(data T) bool {
	dll.lock()
	defer dll.unlock()
	return dll.prepend(data)
}

// Pop removes and returns the last element in the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Pop() (T, bool) {
	dll.lock()
	defer dll.unlock()
	return dll.pop()
}

// PopFirst removes and returns the first element in the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) PopFirst() (T, bool) {
	dll.lock()
	defer dll.unlock()
	return dll.popFirst()
}

// Get returns the element at the specified index in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Get(index int) (T, bool) {
	dll.rLock()
	defer dll.rUnlock()
	return dll.get(index)
}

// Set sets the element at the specified index in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Set(index int, data T) bool {
	dll.lock()
	defer dll.unlock()
	return dll.set(index, data)
}

// Insert adds an element at the specified index in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Insert(index int, data T) bool {
	dll.lock()
	defer dll.unlock()
	return dll.insert(index, data)
}

// Delete removes the element at the specified index in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Delete(index int) bool {
	dll.lock()
	defer dll.unlock()
	return dll.delete(index)
}

// IsEmpty checks if the list is empty in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	dll.rLock()
	defer dll.rUnlock()
	return dll.isEmpty()
}

// Length returns the length of the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Length() int {
	dll.rLock()
	defer dll.rUnlock()
	return dll.length_internal()
}

// Range returns a channel that iterates over the elements in the list in a concurrency-safe manner.
func (dll *DoublyLinkedList[T]) Range() <-chan T {
	dll.rLock()
	ch := make(chan T)
	go func() {
		defer dll.rUnlock()
		node := dll.head
		for node != nil {
			ch <- node.data
			node = node.next
		}
		close(ch)
	}()
	return ch
}

func (dll *DoublyLinkedList[T]) GetNode(index int) *DllNode[T] {
	dll.rLock()
	defer dll.rUnlock()
	return dll.getNode(index)
}

func (dll *DoublyLinkedList[T]) DeleteNode(node *DllNode[T]) bool {
	dll.lock()
	defer dll.unlock()
	return dll.deleteNode(node)
}

func (dll *DoublyLinkedList[T]) AppendNode(node *DllNode[T]) bool {
	dll.lock()
	defer dll.unlock()
	return dll.appendNode(node)
}

func (dll *DoublyLinkedList[T]) PrependNode(node *DllNode[T]) bool {
	dll.lock()
	defer dll.unlock()
	return dll.prependNode(node)
}
