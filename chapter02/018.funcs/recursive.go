package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func guessNumber(left int, right int) {
	guessed := (left + right) / 2

	var getFromInput string
	fmt.Println("我猜是：", guessed)
	fmt.Println("如果高了， 输入1；如果低了，输入0；对了，输入9")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guessNumber(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是改变主意了")
			return
		}
		guessNumber(guessed+1, right)
	case "9":
		fmt.Println("你心里想的数字是：", guessed)
	}
}
