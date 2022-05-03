package main

import "fmt"

func printArray(myArray [4]int)  {
	// 对于固定长度的数组传参是传值
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}

	myArray[0] = 100	// 静态语言特性，传参传的是值
}
func main() {
	// 固定长度的数组
	var myArrray1 [10]int

	myArray2 := [10]int{1,2,3,4}
	myArray3 := [4]int{11,22,33,44}

	for i := 10; i < len(myArrray1); i++ {
		fmt.Println(myArrray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	fmt.Printf("myArray1 type is %T\n", myArrray1)
	fmt.Printf("myArray2 type is %T\n", myArray3)

	printArray(myArray3)
}
