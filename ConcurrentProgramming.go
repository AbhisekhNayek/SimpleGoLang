package main

import (
	"fmt"
	"sync"
	"time"
)

// Function that squares a number and sends the result to a channel
func square(number int, resultChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second) // Simulate some work
	squared := number * number
	resultChannel <- squared
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChannel := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, resultChannel, &wg)
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Read from the channel and print results
	for squared := range resultChannel {
		fmt.Printf("Squared: %d\n", squared)
	}
}
