package main

import (
	"encoding/json"
	"fmt"
)

// Person represents a person with a name and age
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create a Person struct
	person := Person{Name: "John Doe", Age: 30}

	// Encode the Person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	fmt.Printf("Encoded JSON: %s\n", jsonData)

	// Decode the JSON back to a Person struct
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Decoded Person: %+v\n", decodedPerson)
}
