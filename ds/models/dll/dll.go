// Description: This package contains the implementation of a doubly linked list.

package dll

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
}

func ZeroValue[T any]() T {
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

// Append adds an element to the end of the list.
func (dll *DoublyLinkedList[T]) Append(data T) bool {
	node := &DllNode[T]{data: data}
	return dll.AppendNode(node)
}

// AppendNode adds a node to the end of the list.
func (dll *DoublyLinkedList[T]) AppendNode(node *DllNode[T]) bool {
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

// Prepend adds a node to the beginning of the list.
func (dll *DoublyLinkedList[T]) Prepend(data T) bool {
	node := &DllNode[T]{data: data}
	return dll.PrependNode(node)
}

// PrependNode adds a node to the beginning of the list.
func (dll *DoublyLinkedList[T]) PrependNode(node *DllNode[T]) bool {
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

// Pop removes and returns the last element in the list.
func (dll *DoublyLinkedList[T]) Pop() (T, bool) {
	if node, ok := dll.PopNode(); ok {
		return node.data, true
	} else {
		return ZeroValue[T](), false
	}
}

// PopNode removes and returns the last node in the list.
func (dll *DoublyLinkedList[T]) PopNode() (*DllNode[T], bool) {
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

// PopFirst removes and returns the first element in the list.
func (dll *DoublyLinkedList[T]) PopFirst() (T, bool) {
	if node, ok := dll.PopFirstNode(); ok {
		return node.data, true
	} else {
		return ZeroValue[T](), false
	}

}

// PopFirstNode removes and returns the first element in the list.
func (dll *DoublyLinkedList[T]) PopFirstNode() (*DllNode[T], bool) {
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

// GetNode returns the node at the specified index.
func (dll *DoublyLinkedList[T]) GetNode(index int) *DllNode[T] {
	if index < 0 || index >= dll.length {
		return nil
	}
	node := dll.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// Get returns the element at the specified index.
func (dll *DoublyLinkedList[T]) Get(index int) (T, bool) {
	node := dll.GetNode(index)
	if node == nil {
		return ZeroValue[T](), false
	}
	return node.data, true
}

// Set sets the element at the specified index.
func (dll *DoublyLinkedList[T]) Set(index int, data T) bool {
	node := dll.GetNode(index)
	if node == nil {
		return false
	}
	node.data = data
	return true
}

// Insert adds an element at the specified index.
func (dll *DoublyLinkedList[T]) Insert(index int, data T) bool {
	if index < 0 || index > dll.length {
		return false
	}
	if index == 0 {
		return dll.Prepend(data)
	}
	if index == dll.length {
		return dll.Append(data)
	}
	node := &DllNode[T]{data: data}
	prevNode := dll.GetNode(index - 1)
	nextNode := prevNode.next
	prevNode.next = node
	node.prev = prevNode
	node.next = nextNode
	nextNode.prev = node
	dll.length++
	return true
}

// Delete removes the element at the specified index.
func (dll *DoublyLinkedList[T]) Delete(index int) bool {
	node := dll.GetNode(index)
	if node == nil {
		return false
	}
	return dll.DeleteNode(node)
}

func (dll *DoublyLinkedList[T]) DeleteNode(node *DllNode[T]) bool {
	if node == nil {
		return false
	}
	if node != dll.head && node != dll.tail && node.prev == nil && node.next == nil {
		return false
	}
	if node == dll.head {
		_, ok := dll.PopFirst()
		return ok
	}
	if node == dll.tail {
		_, ok := dll.Pop()
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

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.length == 0
}

func (dll *DoublyLinkedList[T]) Length() int {
	return dll.length
}

// Range returns a channel that iterates over the elements in the list.
func (dll *DoublyLinkedList[T]) Range() <-chan T {
	ch := make(chan T)
	go func() {
		node := dll.head
		for node != nil {
			ch <- node.data
			node = node.next
		}
		close(ch)
	}()
	return ch
}
