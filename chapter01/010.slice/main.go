package main

import "fmt"

func main() {
	fmt.Println("hello===")
	a := "hello"
	fmt.Println(a)
	aBytes := []byte(a) // []int 不能转，[]byte 可以和string互转
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

	fmt.Println("您好===")
	b := "您好"
	fmt.Println(b)
	bBytes := []byte(b)
	bRunes := []rune(b)

	//fmt.Println("byte b：", bBytes, len(bBytes))
	//fmt.Println("rune b：", bRunes, len(bRunes))

	fmt.Println("bBytes：", bBytes, len(bBytes))
	fmt.Println("bRunes：", bRunes, len(bRunes))
	fmt.Println("修改切片内的内容")
	bBytes[0] = 'H'
	bRunes[0] = 'H'
	b = string(bBytes)
	fmt.Println(b)
	c := string(bRunes)
	fmt.Println(c)
}
