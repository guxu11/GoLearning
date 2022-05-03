package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Actor []string `json:"actors"`
}


func main() {
	movie := Movie{"taxi", 1986, []string{"guxu"}}

	// 编码：将结构体转化为json
	jsomStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("error:", err)
		return
	} else {
		fmt.Printf("jsonStr = %s\n", jsomStr)
	}

	// 解码： 将json转化为结构体
	// jsonStr = {"title":"taxi","year":1986,"actors":["guxu"]}
	myMovie := Movie{}
	err2 := json.Unmarshal(jsomStr, &myMovie)
	if  err2!= nil {
		fmt.Println("error: ", err2)
		return
	} else {
		fmt.Printf("%v\n", myMovie)
	}

}