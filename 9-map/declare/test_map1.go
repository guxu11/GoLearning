package main

import "fmt"

func main() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// method2
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	// method3
	myMap3 := map[string] string {
		"one": "php",
		"two": "c++",
		"three": "go",
	}
	fmt.Println(myMap3)

}
