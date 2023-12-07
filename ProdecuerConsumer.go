package main

import (
	"fmt"
	"sync"
)

const bufferSize = 5

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		value, ok := <-ch
		if !ok {
			break // Channel closed, exit the loop
		}
		fmt.Printf("Consumed: %d\n", value)
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a channel with a buffer size
	ch := make(chan int, bufferSize)

	// Start the producer and consumer goroutines
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// Wait for the producer and consumer to finish
	wg.Wait()
}
