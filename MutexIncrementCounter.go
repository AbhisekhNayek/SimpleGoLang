package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock the mutex before accessing the shared counter
	mutex.Lock()
	counter++
	// Unlock the mutex after modifying the counter
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// Create 100 goroutines to increment the counter concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementCounter(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
