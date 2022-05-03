package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// f: pair<type: *os.File, value: "/dev/tty"文件描述符>
	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	fmt.Printf("f.type is %T\n", f)
	// r: pair<type: , value: >
	var r io.Reader
	// f: pair<type: *os.File, value: "/dev/tty"文件描述符>
	r = f

	// r: pair<type: , value: >
	var w io.Writer
	// f: pair<type: *os.File, value: "/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("hello this is a tty\n"))

}
