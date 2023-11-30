package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	num := 5 // You can change this to any non-negative integer
	result := factorial(num)
	fmt.Printf("Factorial of %d is: %d\n", num, result)
}
