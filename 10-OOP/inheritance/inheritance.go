package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (this *Human) Print() {
	fmt.Println("name", this.Name)
	fmt.Println("age", this.Age)
	fmt.Println("gender", this.Gender)
}

func (this *Human) Eat() {
	fmt.Println("Human is eating..")
}

type Superman struct {
	Human

	Level int
}

func (this *Superman) Eat() {
	fmt.Println("Superman is eating..")
}

func (this *Superman) Print() {
	fmt.Println("name", this.Name)
	fmt.Println("age", this.Age)
	fmt.Println("gender", this.Gender)
	fmt.Println("level", this.Level)
}

func main() {
	human := Human{"guxu", 20, "male"}
	human.Print()
	human.Eat()

	fmt.Println("--------")
	sm := Superman{Human{"ziyi", 12, "female"}, 88}
	//var sm Superman
	//sm.Name = "ziyi"
	//sm.Age = 18
	//sm.Gender = "female"
	//sm.Level = 99
	sm.Eat()
	sm.Print()
}
