package main

import (
	//"firstProj5-init/lib1"
	_ "firstProj/5-init/lib1"
	. "firstProj/5-init/lib2"
)

// 包的前面加一个下划线，表示导这个包但不使用，会调用init方法
// 包的前面加一个. 相当于Python的 from pkg import x, 直接使用x
func main() {
	//lib1.Lib1Test()
	Lib2Test()
}
