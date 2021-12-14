package main

import "fmt"

func main() {
	var rmb int
	var busy bool = false

	switch {
	case rmb >= 0 && rmb <= 20:
		fmt.Println("点个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
		fallthrough // 直接并入下一个处理分支而无需判断条件
	case rmb > 20 && rmb <= 200:
		fmt.Println("下个馆子")
	case rmb > 200 && rmb <= 2000:
		fmt.Println("去米其林")
	default:
		fmt.Println("容我想想")
	}
	fmt.Printf("end")
}
