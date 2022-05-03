package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)

	// 区分此时引用的底层数据类型
	value, ok := arg.(Book)
	//fmt.Println("arg is string , value = ", value)

	if !ok {
		fmt.Println("arg is not a Book")
		fmt.Printf("value type is %T\n", arg)
	} else {
		fmt.Println("arg is Book , value = ", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"guxu"}
	myFunc(book)
	myFunc("1")
}
