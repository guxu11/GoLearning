package main

import "fmt"

// type声明一种新的数据类型
type myint int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	// 传递的时book的副本
	book.author = "ziyi"
}

func changeBook2(book *Book) {
	book.author = "ziyi"
}
func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)

	fmt.Println("-------------------------")

	var book1 Book
	book1.title = "Golang"
	book1.author = "guxu"
	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
