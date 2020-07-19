package main

import (
	"fmt"
)

func main() {
	// Declaring variables with explicit types
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	// Declaring variables by implicitly inferring the types
	firstName := "Jack"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	// Useful complex type in Go
	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
