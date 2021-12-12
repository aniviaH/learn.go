package main

import "fmt"

func main() {
	hello()
	helloToSomeone("李四")
	helloToSomeone("张三")

	msg := constructHelloMessage("王五")
	fmt.Println(msg)
}
func hello() {
	fmt.Println("你好，Hello")
}

func helloToSomeone(name string) {
	fmt.Println("你好，", name)
}

func constructHelloMessage(name string) (msg string) {
	//fmt.Println("你好，", name)
	return "你好，" + name
}
