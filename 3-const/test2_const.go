package main

import (
	"fmt"
)

// const 定义枚举类型
const (
	// 可以在const()中加一个关键字 iota，每行的iota都会加1，第一行iota默认为0
	BEIJING = 10 * iota // iota = 0
	SHANGHAI
	NANJING
)

// iota 只有在const中使用
const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 0 + 1, b = 0 + 2
	c, d                      // iota = 1, c = 1 + 1, d = 1 + 2
	e, f = iota + 2, iota + 3 // iota = 2, e = 2 + 2, f = 2 + 3
	i, k                      // iota = 3, i = 3 + 2, f = 3 + 3
)

func main() {
	// 常量（只读属性）
	const length int = 10
	const l = 100
	// length = 100 是不可以的
	//var _ int

	fmt.Println("length = ", length)

	fmt.Println("BEIJING:", BEIJING)
	fmt.Println("SHANGHAI:", SHANGHAI)
	fmt.Println("NANJING", NANJING)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(i, k)
}
