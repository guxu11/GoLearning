package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体的类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}
func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

type NewBook struct {
}

func (this *NewBook) ReadBook() {
	fmt.Println("read a newbook")
}
func (this *NewBook) WriteBook() {
	fmt.Println("write a newbook")
}

func main() {
	// b: pair<type: *Book, value:Book{}地址>
	b := &Book{}
	nb := &NewBook{}

	var r Reader
	r = b
	// r: pair<type: *Book, value:Book{}地址>
	r.ReadBook()

	//var w Writer
	w := nb
	fmt.Printf("type of w is %T\n", w)
	//fmt.Printf("type of err is %T\n", err)
	//fmt.Println(value, err)
	// w: pair<type: *Book, value:Book{}地址>
	w.WriteBook() // 因为w和r具体的type是一致的，所以能断言成功
}
