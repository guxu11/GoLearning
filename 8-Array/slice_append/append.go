package main

import "fmt"

func main() {
	s0 := make([]int, 3)
	fmt.Printf("len = %d, capacity = %d, slice1 = %v\n", len(s0), cap(s0), s0)

	s := make([]int, 3, 5)

	fmt.Printf("len = %d, capacity = %d, slice1 = %v\n", len(s), cap(s), s)

	s = append(s, 1)
	fmt.Printf("len = %d, capacity = %d, slice1 = %v\n", len(s), cap(s), s)

	s = append(s, 2)
	s = append(s, 3)
	fmt.Printf("len = %d, capacity = %d, slice1 = %v", len(s), cap(s), s)

}
