package main

import "fmt"

func printSlice(myArray[]int)  {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100	// slice 作为参数，传的是地址，类似Python
}
func main() {
	mySlice := []int{1,2,3,4}

	printSlice(mySlice)
	fmt.Println("===")

	for _, value := range mySlice {
		fmt.Println("value = ", value)
	}

}
