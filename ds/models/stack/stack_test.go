package stack_test

import (
	"testing"

	"github.com/mmygods/gods/ds/collections"
	"github.com/mmygods/gods/ds/models/stack"
)

func TestStackPushPeekPopEmpty(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
	}{
		{
			name:     "Stack with Elements",
			elements: []int{1, 2, 3},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack.New[int]()

			// Push elements onto the stack
			for _, element := range tt.elements {
				s.Push(element)
			}

			// Verify the length of the stack
			if s.Length() != len(tt.elements) {
				t.Errorf("Expected length %d, but got %d", len(tt.elements), s.Length())
			}

			// Pop elements from the stack and verify the order
			for i := len(tt.elements) - 1; i >= 0; i-- {
				element, _ := s.Peek()
				expected := tt.elements[i]
				if element != expected {
					t.Errorf("Expected %d, but got %d", expected, element)

				}

				element, _ = s.Pop()
				expected = tt.elements[i]

				// Verify the length of the stack
				if s.Length() != i {
					t.Errorf("Expected length %d, but got %d", i, s.Length())
				}

				if element != expected {
					t.Errorf("Expected %d, but got %d", expected, element)
				}
			}

			// Verify that the stack is empty
			if !s.IsEmpty() {
				t.Error("Expected stack to be empty when all elements are popped, but got non-empty stack")
			}

		})
	}
}

func TestStackPopPeekEmpty(t *testing.T) {
	s := stack.New[int]()
	_, ok := s.Pop()
	if ok {
		t.Error("Expected empty stack to return false on pop, but got true")
	}
	_, ok = s.Peek()
	if ok {
		t.Error("Expected empty stack to return false on peek, but got true")
	}

}

func TestStackInterface(t *testing.T) {
	var s collections.Stack[int] = stack.New[int]()
	tests := []struct {
		name       string
		elements   []int
		operations []func(collections.Stack[int])
	}{
		{
			name:     "Stack with Elements",
			elements: []int{1, 2, 3},
			operations: []func(collections.Stack[int]){
				func(s collections.Stack[int]) {
					s.Push(4)
				},
				func(s collections.Stack[int]) {
					s.Pop()
				},
				func(s collections.Stack[int]) {
					s.Peek()
				},
				func(s collections.Stack[int]) {
					s.IsEmpty()
				},
				func(s collections.Stack[int]) {
					s.Length()
				},
			},
		},
		{
			name:     "Empty Stack",
			elements: []int{},
			operations: []func(collections.Stack[int]){
				func(s collections.Stack[int]) {
					s.Pop()
				},
				func(s collections.Stack[int]) {
					s.Peek()
				},
				func(s collections.Stack[int]) {
					s.IsEmpty()
				},
				func(s collections.Stack[int]) {
					s.Length()
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, operation := range tt.operations {
				operation(s)
			}
		})
	}
}
