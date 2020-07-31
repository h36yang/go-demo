package main

import (
	"fmt"
)

// Creating a user struct
type user struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	// Init user object - long form
	var u1 user
	u1.ID = 1
	u1.FirstName = "Bruce"
	u1.LastName = "Wayne"
	fmt.Println(u1)

	// Init user object - short form
	u2 := user{
		ID:        1,
		FirstName: "Bruce",
		LastName:  "Wayne", // need to end with comma because Go automatically appends ";" to end of line otherwise
	}
	fmt.Println(u2)
}
