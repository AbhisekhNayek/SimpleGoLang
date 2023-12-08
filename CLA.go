package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the command-line arguments
	args := os.Args

	fmt.Println("Command-line arguments:")
	for i, arg := range args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
