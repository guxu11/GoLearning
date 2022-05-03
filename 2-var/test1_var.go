package main

import "fmt"

// 声明全局变量
var gA int = 100
var gB = 200

// 不能用 a := 1声明全局变量

func main() {
	// method1
	var a int
	fmt.Println("a = ", a)

	// method 2
	var b int = 100
	fmt.Println(b)

	var bb string = "abcd"
	fmt.Printf("bb type: %T\n", bb)

	// method3 省去数据类型
	var c = 100
	fmt.Println(c)

	// method4 (最常用) 省略var
	d := 100
	fmt.Println("e = ", d)
	fmt.Printf("e type is %T\n", d)

	fmt.Println("gA = ", gA, "gB = ", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	var kk, ll = 100, "ab"
	fmt.Println(xx, yy)
	fmt.Println(kk, ll)

	cd, cb := 100, 200.0
	fmt.Printf("cd = %d, cb = %.2f\n", cd, cb)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println(vv, jj)
}
