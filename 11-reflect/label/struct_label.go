package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"guxu" doc:"mingzi"`
	Age int `info:"100"`
}

func getLabel(str interface{})  {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag_info := t.Field(i).Tag.Get("info")
		fmt.Println("info is ", tag_info)
	}
}
func main() {
	r := resume{"guxu", 20}
	getLabel(&r)
}
