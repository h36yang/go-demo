package main

import (
	"errors"
	"fmt"
)

func main() {
	// Invoking function 1
	startWebServer1()

	// Invoking function 2
	port := 3000
	startWebServer2(port, 2)

	// Invoking function 3
	err := startWebServer3(port)
	fmt.Println(err)

	// Invoking function 4
	x1, x2 := startWebServer4(port)
	fmt.Println(x1, x2)
	// Discarding one of the return values with _
	_, x3 := startWebServer4(port)
	fmt.Println(x3)
}

// Function without parameters
func startWebServer1() {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started")
}

// Function with parameters
// NOTE: Go assumes parameters are of the same type as the last one if not specified
func startWebServer2(port, numberOfRetries int) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries:", numberOfRetries)
}

// Function with return value
// NOTE: Popular paradigm in Go to return an "error" type instead of throwing exceptions
func startWebServer3(port int) error {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return errors.New("Something went wrong")
}

// Function with multiple return values
func startWebServer4(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return 1, errors.New("Something went wrong")
}
