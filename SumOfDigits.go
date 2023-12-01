package main

import (
	"fmt"
)

// Function to calculate the sum of digits in a number
func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := sumOfDigits(num)
	fmt.Printf("Sum of digits in %d is: %d\n", num, result)
}
