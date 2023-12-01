package main

import (
	"fmt"
	"math"
)

// Function to check if a number is prime
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}
