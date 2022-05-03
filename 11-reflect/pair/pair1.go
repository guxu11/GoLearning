package main

import "fmt"

func main() {
	var a string
	// pair<statictype: string, value: "aceld">
	a = "aceld"

	// pair<type:string, value: "aceld">
	// 把type指针指向string，value指针指向aceld
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
