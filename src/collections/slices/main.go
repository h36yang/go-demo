package main

import (
	"fmt"
)

func main() {
	// Init slice - array without fixed size
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	s1 = append(s1, 4, 5, 6, 7)
	fmt.Println(s1)

	// Create slices from slice
	s2 := s1[1:]  // from index 1 to the end
	s3 := s1[:3]  // from beginning to (not including) index 3
	s4 := s1[1:3] // from index 1 to (not including) index 3
	fmt.Println(s2, s3, s4)
}
