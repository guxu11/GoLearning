package main

import (
	"fmt"
	"time"
)

func route1() {
	defer fmt.Println("first routine ended")

	for i := 0; i < 3; i++ {
		fmt.Println("routine 1 ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	defer fmt.Println("main routine end")
	go route1()

	go func() {
		defer fmt.Println("secpnd routine end")
		for i := 0; i < 3; i++ {
			fmt.Println("routine 2 ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("main routine ", i)
		time.Sleep(1 * time.Second)
	}
}
