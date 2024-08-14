package deque_test

import (
	"testing"

	"github.com/mmygods/gods/ds/collections"
	"github.com/mmygods/gods/ds/models/dll"
)

func TestDequeInterface(t *testing.T) {
	var deque collections.Deque[int] = &dll.DoublyLinkedList[int]{}
	deque.Append(1)
	deque.Append(2)
	deque.Append(3)
	deque.Prepend(4)
	deque.Prepend(5)

	if deque.Length() != 5 {
		t.Errorf("Expected length to be 5, got %d", deque.Length())
	}

	if data, ok := deque.Pop(); !ok || data != 3 {
		t.Errorf("Expected 3, got %d", data)
	}

	if data, ok := deque.PopFirst(); !ok || data != 5 {
		t.Errorf("Expected 5, got %d", data)
	}

	if deque.Length() != 3 {
		t.Errorf("Expected length to be 3, got %d", deque.Length())
	}

	if data, ok := deque.Pop(); !ok || data != 2 {
		t.Errorf("Expected 2, got %d", data)
	}

	if data, ok := deque.PopFirst(); !ok || data != 4 {
		t.Errorf("Expected 4, got %d", data)
	}

	if data, ok := deque.Pop(); !ok || data != 1 {
		t.Errorf("Expected 1, got %d", data)
	}

	if deque.Length() != 0 {
		t.Errorf("Expected length to be 0, got %d", deque.Length())
	}

	if _, ok := deque.Pop(); ok {
		t.Error("Expected ok to be false for empty deque")
	}

	if _, ok := deque.PopFirst(); ok {
		t.Error("Expected ok to be false for empty deque")
	}

	if !deque.IsEmpty() {
		t.Error("Expected deque to be empty")
	}
}
