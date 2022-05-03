package main

import "fmt"

func main() {
	// method1
	//slice1 := []int{1,2,3}

	// method2
	//var slice1 int	// 声明slice1是一个切片，但是没有分配空间
	//slice1 = make([]int, 3)	// 开辟三个空间，默认值是1

	// method3	将第二种方法一步表示
	//var slice1 []int = make([]int, 3)

	// method4	最常用
	slice1 := make([]int, 3)
	fmt.Println("len = %d, slice = %v", len(slice1), slice1)

	// 判断一个slice是否为空切片
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
