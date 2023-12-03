package main

import "fmt"

// Function to perform bubble sort on a slice of integers
func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Printf("Unsorted array: %v\n", arr)

	bubbleSort(arr)

	fmt.Printf("Sorted array: %v\n", arr)
}
