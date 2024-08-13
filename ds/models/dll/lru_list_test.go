package dll_test

import (
	"testing"

	"github.com/mmygods/gods/ds/models/dll"
)

func TestDoublyLinkedListDeleteNonExistingAndExistingNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		index    int
		expected bool
	}{
		{
			name:     "Delete non-existing node",
			elements: []int{1, 2, 3},
			index:    3,
			expected: false,
		},
		{
			name:     "Delete node from empty list",
			elements: []int{},
			index:    0,
			expected: false,
		},
		{
			name:     "Delete existing node",
			elements: []int{1, 2, 3},
			index:    1,
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			if list.Delete(test.index) != test.expected {
				t.Errorf("Should return %t when deleting a node", test.expected)
			}
		})
	}
}

func TestDoublyLinkedListDeleteExistingNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		index    int
		expected []int
	}{
		{
			name:     "Delete middle node",
			elements: []int{1, 2, 3},
			index:    1,
			expected: []int{1, 3},
		},
		{
			name:     "Delete first node",
			elements: []int{1, 2, 3},
			index:    0,
			expected: []int{2, 3},
		},
		{
			name:     "Delete last node",
			elements: []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			if !list.DeleteNode(list.GetNode(test.index)) {
				t.Errorf("Should return true when deleting a node")
			}

			if list.Length() != len(test.expected) {
				t.Errorf("Length should be %d after deleting a node", len(test.expected))
			}

			for i, expected := range test.expected {
				if data, _ := list.Get(i); data != expected {
					t.Errorf("Data should be %d after deleting a node", expected)
				}
			}
		})
	}
}

func TestDoublyLinkedListAppendNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		node     *dll.DllNode[int]
		expected []int
	}{
		{
			name:     "Append node to empty list",
			elements: []int{},
			node:     dll.NewNode(1),
			expected: []int{1},
		},
		{
			name:     "Append node to non-empty list",
			elements: []int{1, 2},
			node:     dll.NewNode(3),
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			list.AppendNode(test.node)

			if list.Length() != len(test.expected) {
				t.Errorf("Length should be %d after appending a node", len(test.expected))
			}

			for i, expected := range test.expected {
				if data, _ := list.Get(i); data != expected {
					t.Errorf("Data should be %d after appending a node", expected)
				}
			}
		})
	}
}

func TestDoublyLinkedListPrependNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		node     *dll.DllNode[int]
		expected []int
	}{
		{
			name:     "Prepend node to empty list",
			elements: []int{},
			node:     dll.NewNode(1),
			expected: []int{1},
		},
		{
			name:     "Prepend node to non-empty list",
			elements: []int{1, 2},
			node:     dll.NewNode(3),
			expected: []int{3, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			list.PrependNode(test.node)

			if list.Length() != len(test.expected) {
				t.Errorf("Length should be %d after prepending a node", len(test.expected))
			}

			for i, expected := range test.expected {
				if data, _ := list.Get(i); data != expected {
					t.Errorf("Data should be %d after prepending a node", expected)
				}
			}
		})
	}
}

func TestDoublyLinkedListDeleteNodeWithTailNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		node     *dll.DllNode[int]
		expected []int
	}{
		{
			name:     "Delete tail node",
			elements: []int{1, 2, 3},
			node:     dll.NewNode(3),
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			if list.DeleteNode(test.node) {
				t.Errorf("Should return false when deleting a node that does not belong to the list")
			}
			tailNode := list.GetNode(list.Length() - 1)
			if !list.DeleteNode(tailNode) {
				t.Errorf("Should return true when deleting the tail node")
			}
			if list.Length() != len(test.expected) {
				t.Errorf("Length should be %d after deleting the tail node", len(test.expected))
			}

			for i, expected := range test.expected {
				if data, _ := list.Get(i); data != expected {
					t.Errorf("Data should be %d after deleting the tail node", expected)
				}
			}
		})
	}
}

func TestDoublyLinkedListDeleteNodeWithHeadNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		node     *dll.DllNode[int]
		expected []int
	}{
		{
			name:     "Delete head node",
			elements: []int{1, 2, 3},
			node:     dll.NewNode(1),
			expected: []int{2, 3},
		},
		{
			name:     "Delete nil node",
			elements: []int{1, 2, 3},
			node:     nil,
			expected: []int{2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, element := range test.elements {
				list.Append(element)
			}

			if list.DeleteNode(test.node) {
				t.Errorf("Should return false when deleting a node that does not belong to the list")
			}
			headNode := list.GetNode(0)
			if !list.DeleteNode(headNode) {
				t.Errorf("Should return true when deleting the head node")
			}

			if list.Length() != len(test.expected) {
				t.Errorf("Length should be %d after deleting the head node", len(test.expected))
			}

			var node = list.GetNode(0)
			if node == nil {
				t.Errorf("Should return the next node when getting the head node")
				return
			}

			if node.GetData() != test.expected[0] {
				t.Errorf("Data should be %d after deleting the head node", test.expected[0])
			}
		})
	}
}
