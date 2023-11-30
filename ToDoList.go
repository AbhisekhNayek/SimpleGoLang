package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	todoList := []string{}

	for {
		fmt.Print("Enter a task (or 'exit' to quit): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		task := scanner.Text()

		if task == "exit" {
			break
		}

		todoList = append(todoList, task)
	}

	fmt.Println("\nYour Todo List:")
	for i, task := range todoList {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
