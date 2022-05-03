package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}

	s1 := s[0:3]

	fmt.Println("s1 = ", s1)

	s1[0] = 20	// s1和s地址一样
	fmt.Println("s1 = ", s1)
	fmt.Println("s = ", s)

	// copy
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println("s2 = ", s2)
}
