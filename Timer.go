package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 5 * time.Second

	timer := time.NewTimer(duration)

	fmt.Printf("Timer set for %s. Waiting...\n", duration)

	// Wait for the timer to expire
	<-timer.C

	fmt.Println("Timer expired! Time to perform some action.")
}
