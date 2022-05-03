package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 多个返回值， 匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 66, 77

}

// 多个返回值， 有形参名
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("r1 = ", r1, ", r2 = ", r2)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}
func main() {
	c := foo1("a", 100)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("a2", 99)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 100)
	fmt.Println("ret1:", ret1, "ret2:", ret2)

	ret1, ret2 = foo4("foo3", 100)
	fmt.Println("ret1:", ret1, "ret2:", ret2)

}
