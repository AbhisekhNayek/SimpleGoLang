package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Function to read data from a file
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Function to write data to a file
func writeFile(filename string, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	return err
}

func main() {
	// Write data to a file
	err := writeFile("example.txt", "Hello, this is a file I/O example.")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}
	fmt.Println("Data written to file successfully.")

	// Read data from the same file
	content, err := readFile("example.txt")
	if err != nil {
		fmt.Printf("Error reading from file: %s\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", content)
}
