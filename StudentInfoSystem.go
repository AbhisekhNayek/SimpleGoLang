package main

import "fmt"

// Define a struct for student information
type Student struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Create instances of the Student struct
	student1 := Student{ID: 1, Name: "John Doe", Age: 20}
	student2 := Student{ID: 2, Name: "Jane Smith", Age: 22}

	// Display information about the students
	fmt.Printf("Student 1 - ID: %d, Name: %s, Age: %d\n", student1.ID, student1.Name, student1.Age)
	fmt.Printf("Student 2 - ID: %d, Name: %s, Age: %d\n", student2.ID, student2.Name, student2.Age)
}
