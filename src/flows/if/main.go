package main

type user struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := user{
		ID:        1,
		FirstName: "Bruce",
		LastName:  "Wayne",
	}
	u2 := user{
		ID:        2,
		FirstName: "Clark",
		LastName:  "Kent",
	}

	// If/Else block
	if u1.ID == u2.ID {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Different users!")
	}
}
