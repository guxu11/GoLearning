package main

import "fmt"

// 类和方法首字母大写，表示其他包也能访问
type Hero struct {
	// 如果属性首字母大写，表示是共有变量，否则是私有
	Name  string
	Ad    int
	Level int
}

//func (this Hero) Show()  {
//	fmt.Println("Name = ", this.Name)
//	fmt.Println("Ad = ", this.Ad)
//	fmt.Println("Level = ", this.Level)
//}
//
//func (this Hero) GetName() string {
//	return this.Name
//}
//
//func (this Hero) SetName(newName string)  {
//	// this 是调用该方法的对象的一个副本
//	this.Name = newName
//}
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// this 是调用该方法的对象的一个副本
	this.Name = newName
}

func main() {
	hero := Hero{Name: "Zhangsan", Ad: 100, Level: 1}

	hero.SetName("guxu")
	hero.Show()
}
