package main

import "fmt"

func main() {
	var rmb int = 20 // 修改它，看看它会做什么
	// 如果不超过20，点个外卖
	// 如果不超过200，下个馆子
	// 如果不超过2000， 去米其林
	// 如果不超过20000，去新马泰
	// 如果再多，容我想想

	if rmb <= 20 {
		fmt.Println("点外卖")
	} else if rmb <= 200 {
		fmt.Println("下个馆子")
	} else if rmb <= 2000 {
		fmt.Println("去米其林")
	} else if rmb <= 20000 {
		fmt.Println("去新马泰")
	} else {
		fmt.Println("容我想想")
	}

}
