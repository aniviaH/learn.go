package main

import (
	"fmt"
	"runtime/debug"
)

func panicAndRecover() {
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
		debug.PrintStack()
	}
}
