package main

import "fmt"

// Function to generate Fibonacci sequence up to n terms
func generateFibonacci(n int) []int {
	fibonacci := make([]int, n)
	fibonacci[0], fibonacci[1] = 0, 1

	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}

func main() {
	var n int
	fmt.Print("Enter the number of Fibonacci terms to generate: ")
	fmt.Scan(&n)

	fibonacci := generateFibonacci(n)
	fmt.Printf("Fibonacci sequence up to %d terms: %v\n", n, fibonacci)
}
