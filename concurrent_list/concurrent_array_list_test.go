package concurrent_list


import (
	"testing"
	"sync"
	"fmt"
)


// TestConcurrentArryList_Race tests the thread safety under high concurrency.
func TestConcurrentArrayList_Race(t *testing.T) {
	// Create a concurrent list.
	list := NewConcurrentArrayList[int](10)
	// WaitGroup is used wait for all goroutines of finish.
	var wg sync.WaitGroup
	// Number of concurrent operations.
	const NumRoution int = 1000
	// Launch multiple goroutines to append data simulataneously.
	for i := 0; i < NumRoution; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Star a new goroutine(thread).
		go func(val int) {
			// Ensure Done is called when the goroutine finish.
			defer wg.Done()
			list.Append(i)
		}(i)
	}
	// Block until all goroutines have called Done.
	wg.Wait()
	// Check the result.
	if list.Len() != NumRoution {
		t.Errorf("Excepted length %d, but got %d. Data race occurred!",
		NumRoution, list.Len())
	}
	// Log
	fmt.Printf("Successfully appended %d elements concurrently.\n", list.Len())
}