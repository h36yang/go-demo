package main

import (
	"fmt"
)

func main() {
	// Init map - [key]value
	m := map[string]int{"foo": 42, "bar": 27}
	fmt.Println(m)
	fmt.Println(m["foo"])

	// Change value by key
	m["foo"] = 23
	fmt.Println(m)

	// Delete element by key
	delete(m, "bar")
	fmt.Println(m)
}
