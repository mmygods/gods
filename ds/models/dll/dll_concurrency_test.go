package dll

import (
	"sync"
	"testing"
	"time"
)

func TestDLLConcurrency(t *testing.T) {
	// Create a new DLL instance
	dll := &DoublyLinkedList[int]{}

	// Number of routines for appending and popping
	numRoutines := 1000

	// Wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Append routine
	appendRoutine := func() {
		defer wg.Done()

		// Append elements to the DLL
		for i := 0; i < 1000; i++ {
			dll.Append(i)
		}
	}

	// Pop routine
	popRoutine := func() {
		defer wg.Done()

		// Pop elements from the DLL
		for i := 0; i < 1000; i++ {
			for _, ok := dll.Pop(); !ok; _, ok = dll.Pop() {
				time.Sleep(time.Millisecond)
				// Retry until an element is successfully
				// popped from the DLL
			}
		}
	}

	// Start the append routines
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go appendRoutine()
	}
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go popRoutine()
	}

	// Wait for all routines to finish
	wg.Wait()

	// Verify the final state of the DLL
	// Add your verification logic here
	if dll.Length() != 0 {
		t.Errorf("Expected DLL length 0, but got %d", dll.Length())
	}
	// Add assertions or checks to validate the final state of the DLL
	// For example, you can check if the DLL is empty or has the expected number of elements

	// If the assertions fail, you can use t.Errorf to report the failure
	// For example:
	// if dll.Len() != expectedLength {
	//     t.Errorf("Expected DLL length %d, but got %d", expectedLength, dll.Len())
	// }
}
