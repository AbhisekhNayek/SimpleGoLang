package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run add.go <num1> <num2>")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	num2, err2 := strconv.ParseFloat(os.Args[2], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}

	sum := num1 + num2
	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", num1, num2, sum)
}
