package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	//var a rune = 'a'

	// rune 解释型字符串
	s := "I am an interpreted string literal" +
		"\n\tn 111\n222"
	fmt.Println(s)

	// 原始字符串
	s2 := `After a backslash , certain single character escapes represent
	special values
		n is a line feed or new line
		t is a tab
		111
		222`
	fmt.Println(s2)

	// 字符串拼接
	var z int = 1
	var s3 string = " egg"
	var intToString = strconv.Itoa(z)
	var breakfast string = intToString + s3
	fmt.Println(breakfast)

	// 使用缓冲区拼接字符串
	var buf bytes.Buffer
	for i := 0; i < 500; i++ {
		buf.WriteString("z")
	}
	fmt.Println(buf.String())
}
