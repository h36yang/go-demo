package main

func main() {
	// 1. Loop till condition
	var i int
	for i < 5 {
		println("i =", i) // built-in println that is less sophisticated than fmt.Println
		i++
		// Keyword "break"    - breaks out of the loop early
		// Keyword "continue" - skips an iteration
	}

	// 2. Loop till condition with post clause
	for j := 0; j < 5; j++ {
		println("j =", j)
	}

	// 3. Infinite loop
	var k int
	for {
		if k == 5 {
			break
		}
		println("k =", k)
		k++
	}

	// 4. Loop over collections
	// Slice
	slice := []int{1, 2, 3}
	for idx, val := range slice {
		println(idx, val)
	}
	// Map
	ports := map[string]int{"http": 80, "https": 443}
	for key, val := range ports {
		println(key, val)
	}
}
