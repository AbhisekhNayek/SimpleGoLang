package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Worker %d", id)
	}
}

func main() {
	// Create two channels for communication
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines
	go worker(1, ch1)
	go worker(2, ch2)

	// Use select to receive messages from the goroutines
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
