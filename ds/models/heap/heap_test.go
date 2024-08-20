// Description: Test file for heap data structure.
// Since heap is already implemented in Go's standard library, we will test the standard library's heap package.
package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	// Create a new heap
	h := &intHeap{2, 1, 5, 3, 4}

	// Initialize the heap
	heap.Init(h)

	// Push an element to the heap
	heap.Push(h, 6)

	// Pop the top element from the heap
	popped := heap.Pop(h)
	fmt.Println("Popped element:", popped)

	// Verify the heap property
	for i := 0; i < h.Len(); i++ {
		if h.Less(i, (i-1)/2) {
			t.Errorf("Heap property violated at index %d", i)
		}
	}
}
