package main

import (
	"fmt"
)

// Note: in Go, the size of the array goes before the data type of the array elements
func main() {
	// Init array - long form
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	// Init array - short form
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
}
