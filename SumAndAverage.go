package main

import "fmt"

// Function to calculate the sum and average of numbers in a slice
func calculateSumAndAverage(numbers []float64) (float64, float64) {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	return sum, average
}

func main() {
	numbers := []float64{10.5, 20.0, 30.5, 40.0, 50.5}

	sum, average := calculateSumAndAverage(numbers)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", average)
}
