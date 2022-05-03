package main

import (
	"fmt"
)

func printMap(cityMap map[string]string)  {
	// map 作为参数是一个引用传递
	for k, v := range cityMap {
		fmt.Println("key", k)
		fmt.Println("value", v)
	}

	cityMap["Japan"] = "tokyo"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["china"] = "Beijing"
	cityMap["USA"] = "WDC"

	// 遍历
	printMap(cityMap)
	// delete
	//delete(cityMap, "china")

	// change
	//cityMap["USA"] = "DC"

	fmt.Println("----------------------")

	printMap(cityMap)

}