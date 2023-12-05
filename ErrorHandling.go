package main

import (
	"errors"
	"fmt"
)

// Function that returns an error if the input is negative
func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return fmt.Sqrt(n), nil
}

func main() {
	var num float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result, err := squareRoot(num)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Square root of %.2f is: %.2f\n", num, result)
	}
}
