package main

import "fmt"

func main() {
	c := make(chan int, 3)

	go func() {
		fmt.Println(<-c)
		//c <- 1
	}()
}
