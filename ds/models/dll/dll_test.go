package dll_test

import (
	"testing"

	"github.com/mmygods/gods/ds/collections"
	"github.com/mmygods/gods/ds/models/dll"
)

func TestDoublyLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(list collections.List[int])
		expected func(list collections.List[int]) bool
	}{
		{
			name: "Empty List",
			actions: func(list collections.List[int]) {
				// No actions
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 0 && list.IsEmpty()
			},
		},
		{
			name: "Append Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 3 && !list.IsEmpty()
			},
		},
		{
			name: "Prepend Elements",
			actions: func(list collections.List[int]) {
				list.Prepend(1)
				list.Prepend(2)
				list.Prepend(3)
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 3 && !list.IsEmpty()
			},
		},
		{
			name: "Insert Elements",
			actions: func(list collections.List[int]) {
				list.Insert(0, 1)
				list.Insert(1, 2)
				list.Insert(2, 3)
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 3 && !list.IsEmpty()
			},
		},
		{
			name: "Get Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
			},
			expected: func(list collections.List[int]) bool {
				data, _ := list.Get(0)
				if data != 1 {
					return false
				}
				data, _ = list.Get(1)
				if data != 2 {
					return false
				}
				data, _ = list.Get(2)
				if data != 3 {
					return false
				}
				_, ok := list.Get(3)
				return !ok
			},
		},
		{
			name: "Set Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.Set(0, 4)
				list.Set(1, 5)
				list.Set(2, 6)
			},
			expected: func(list collections.List[int]) bool {
				data, _ := list.Get(0)
				if data != 4 {
					return false
				}
				data, _ = list.Get(1)
				if data != 5 {
					return false
				}
				data, _ = list.Get(2)
				if data != 6 {
					return false
				}
				return !list.Set(3, 7)
			},
		},
		{
			name: "Pop Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.Pop()
				list.PopFirst()
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 1 && !list.IsEmpty()
			},
		},
		{
			name: "Delete Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.Delete(1)
				list.Delete(0)
				list.Delete(0)
			},
			expected: func(list collections.List[int]) bool {
				return list.Length() == 0 && list.IsEmpty()
			},
		},
		{
			name: "Range Elements",
			actions: func(list collections.List[int]) {
				list.Append(1)
				list.Append(2)
				list.Append(3)
			},
			expected: func(list collections.List[int]) bool {
				for data := range list.Range() {
					if data == 0 {
						return false
					}
				}
				return true
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var list collections.List[int] = &dll.DoublyLinkedList[int]{}
			test.actions(list)
			if !test.expected(list) {
				t.Errorf("Test failed: %s", test.name)
			}
		})
	}

}

func TestDoublyLinkedListPopEmptyList(t *testing.T) {
	list := &dll.DoublyLinkedList[int]{}
	_, ok := list.Pop()
	if ok {
		t.Error("Test failed: Pop empty list")
	}
	_, ok = list.PopFirst()
	if ok {
		t.Error("Test failed: PopFirst empty list")
	}
}

func TestDoublyLinkedListPopSingletonList(t *testing.T) {
	list := &dll.DoublyLinkedList[int]{}
	list.Append(1)
	data, ok := list.Pop()
	if !ok || data != 1 {
		t.Error("Test failed: Pop singleton list")
	}
}

func TestDoublyLinkedListInsertAtIndex(t *testing.T) {
	tests := []struct {
		name         string
		initial_list []int
		index        int
		data         int
		expected     []int
	}{
		{
			name:         "Insert at head",
			initial_list: []int{1, 2, 3},
			index:        0,
			data:         4,
			expected:     []int{4, 1, 2, 3},
		},
		{

			name:         "Insert at tail",
			initial_list: []int{1, 2, 3},
			index:        3,
			data:         4,
			expected:     []int{1, 2, 3, 4},
		},
		{
			name:         "Insert at middle",
			initial_list: []int{1, 2, 3},
			index:        1,
			data:         4,
			expected:     []int{1, 4, 2, 3},
		},
		{
			name:         "Insert at invalid index",
			initial_list: []int{1, 2, 3},
			index:        4,
			data:         4,
			expected:     []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := &dll.DoublyLinkedList[int]{}
			for _, data := range test.initial_list {
				list.Append(data)
			}
			list.Insert(test.index, test.data)
			for i, data := range test.expected {
				if d, _ := list.Get(i); d != data {
					t.Errorf("Test failed: %s", test.name)
				}
			}
		})
	}
}
