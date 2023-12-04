package main

import (
	"fmt"
	"net/http"
)

// Handler function for the web server
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is a basic web server!")
}

func main() {
	// Register the helloHandler function for the "/hello" route
	http.HandleFunc("/hello", helloHandler)

	// Start the web server on port 8080
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
