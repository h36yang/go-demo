package main

import (
	"fmt"
)

func main() {
	// * in this line is the "pointer" operator, meaning it's a pointer to a string variable
	var firstNamePtr *string = new(string)
	// * in this line is the "de-referencing" operator, which retrieves the variable value from the pointer
	*firstNamePtr = "Bruce"
	fmt.Println(firstNamePtr, *firstNamePtr)

	firstName := "Clark"
	// & is the "address of" operator, which retrieves the address of the variable in the memory
	ptr := &firstName
	fmt.Println(ptr, *ptr)
	// Change the value of the variable, but memory location remains unchanged
	firstName = "Diana"
	fmt.Println(ptr, *ptr)
}
