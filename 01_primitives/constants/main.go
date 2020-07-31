package main

import (
	"fmt"
)

// Constant Block
/*
Since constants are evaluated at compile time,
we can't assign it from a function return value
because functions are evaluated at run time.
*/
const (
	first  = 1        // automatically evaluated as int type
	second = "second" // automatically evaluated as string type
)

// iota - starts from 0 and auto increments
// constant expression - can specify operations when assigning constants
const (
	x1 = iota + 3  // expecting 0 + 3 = 3
	x2 = 5 << iota // expecting 5 << 1 = 10
)

// iota resets in a new constant block
// Constant inherits the same expression from the line above when not specified
const (
	y1 = iota + 2 // expecting 0 + 2 = 2
	y2            // expecting 1 + 2 = 3
	y3            // expecting 2 + 2 = 4
)

func main() {
	// Declare constant without type
	const c1 = 2
	// Dynamically treating c1 as int
	fmt.Println(c1 + 2)
	// Dynamically treating c1 as float32
	fmt.Println(c1 + 1.3)

	// Declare constant with type
	const c2 int = 2
	// int + int works fine
	fmt.Println(c2 + 2)
	// int + float32 doesn't work so have to convert first
	fmt.Println(float32(c2) + 1.3)

	// Constants from package level
	fmt.Println(first, second)
	fmt.Println(x1, x2)
	fmt.Println(y1, y2, y3)
}
