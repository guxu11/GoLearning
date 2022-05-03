package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func (this *User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}
func main() {
	var user User
	user = User{
		Name:   "guxu",
		Age:    18,
		Gender: "male",
	}
	//fmt.Println(user)
	//ChangeValue(user)
	//fmt.Println(user)
	DoFiledAndMethod(user)

}

func ChangeValue(user User) {
	user.Name = "1122"
}

func DoFiledAndMethod(arg interface{}) {
	InputType := reflect.TypeOf(arg)
	fmt.Println("type = ", InputType.Name())
	fmt.Println("Kind = ", InputType.Kind())
	InputValue := reflect.ValueOf(arg)
	fmt.Println("value = ", InputValue)

	// 通过type获取里面的字段
	// 1. 获取interface的reflect.type，通过Type得到NumField，进行遍历
	// 2. 得到每个field，数据类型
	// 3. 通过field有一个Interface()方法，得到相应的value
	for i := 0; i < InputType.NumField(); i++ {
		field := InputType.Field(i)
		value := InputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//InputTypePointer := &InputType
	fmt.Println("nummethod = ", InputType.NumMethod())

	// 通过type获得里面的方法
	for i := 0; i < InputType.NumMethod(); i++ {
		m := InputType.Method(i)
		fmt.Printf("%s: %v", m.Name, m.Type)
	}
	//method, _ := InputType.MethodByName("Call")
	//fmt.Println("method type is ", reflect.TypeOf(method))
	//fmt.Printf("method is %s, type is %s, kind is %s.\n", method.Name, method.Type, method.Type.Kind())
}
