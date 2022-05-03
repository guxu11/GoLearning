package main

import (
	"fmt"
)

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 继承接口只要实现其中的方法
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleeping..")
}
func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleeping")
}
func (this *Dog) GetColor() string {
	return this.Color
}
func (this *Dog) GetType() string {
	return "dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}
func main() {

	var animal AnimalIF = &Cat{"Green"}
	animal.Sleep()
	fmt.Printf("animal type is: %T\n", animal)

	animal = &Dog{"yellow"}
	animal.Sleep()
	fmt.Printf("animal type is: %T\n", animal)

	fmt.Println("-----------------")

	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat)
	showAnimal(&dog)
}
