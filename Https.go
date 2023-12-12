package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchData(url string, ch chan string) {
	// Make an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching data from %s: %v", url, err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}

	// Send the fetched data to the channel
	ch <- string(data)
}

func main() {
	// URL to fetch data from
	url := "https://www.example.com"

	// Create a channel to receive the fetched data or error messages
	dataChannel := make(chan string)

	// Start a goroutine to fetch data
	go fetchData(url, dataChannel)

	// Wait for the fetched data or error message
	result := <-dataChannel

	fmt.Println(result)
}
