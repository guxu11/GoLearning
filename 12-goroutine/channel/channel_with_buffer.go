package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子goroutine正在运行, 发送的元素 = ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))

		}
	}()

	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		num := <-c
		fmt.Println("num = ", num)
	}

	defer fmt.Println("main goroutine结束")
}
