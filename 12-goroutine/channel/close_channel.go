package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("writing ", i)
		}
		close(c)
		fmt.Println("channel closed!")
	}()

	for {
		time.Sleep(1 * time.Second)
		if data, ok := <-c; ok {
			fmt.Println("data = ", data)
		} else {
			break
		}
	}

	defer fmt.Println("main routine end")
}
