package main

import "fmt"

// Function to perform linear search in a slice
func linearSearch(arr []int, target int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var target int

	fmt.Print("Enter a number to search in the slice: ")
	fmt.Scan(&target)

	if linearSearch(arr, target) {
		fmt.Printf("%d is present in the slice.\n", target)
	} else {
		fmt.Printf("%d is not present in the slice.\n", target)
	}
}
