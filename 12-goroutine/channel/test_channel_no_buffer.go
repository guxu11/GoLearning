package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行")

		c <- 666
	}()

	num := <-c

	fmt.Println("num = ", num)
	defer fmt.Println("main routine end")
}
