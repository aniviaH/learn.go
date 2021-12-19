package main

import (
	"fmt"
)

func panicAndRecover() {
	fmt.Println("main.go中的全局变量tall, weight: ", tall, weight)

	defer coverPanic()
	defer coverPanicUpgraded() // 这样是能抓住严重错误的
	//defer func() { // 抓不住严重错误
	//	coverPanicUpgraded()
	//}()

	var nameScore map[string]int = nil
	nameScore["张三"] = 90
}

func coverPanic() { // 未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出了严重故障：", r)
		}
	}()
}

func coverPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("系统出了严重故障：", r)
		//debug.PrintStack()
	}
}
