package main

import "fmt"

func fun1()  {
	fmt.Println("A")
}
func fun2()  {
	fmt.Println("B")
}
func fun3()  {
	fmt.Println("C")
}
func main() {
	defer fun1()	// 先入栈，最后执行
	defer fun2()
	defer fun3()

}
