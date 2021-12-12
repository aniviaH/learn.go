package main

import "fmt"

var name string = "GO-全局"

func main() {
	//defer func() { // 如果该函数放在在main中name的定义之前，同时全局的name又不定义，则直接编译报错 Unresolved reference name
	//	fmt.Println("main中定义的匿名函数访问name---3", name1)
	//}()

	//name := "GO-main中"
	defer deferFunc()
	defer fmt.Println("defer println访问name---2", name)

	defer func() {
		fmt.Println("main中定义的匿名函数访问name---3", name)
	}()

	//var name string = "python"
	name := "GO-main中"
	name = "python"
	fmt.Println("main中访问name", name) // python
	name = "java"

	defer func() {
		fmt.Println("main中定义的匿名函数访问name---4", name)
	}()

	defer func(name string) {
		fmt.Println("main中定义的匿名函数并带参数访问name---5", name)
	}(name)
}
func deferFunc() {
	fmt.Println("deferFuc访问name---1", name)
}
