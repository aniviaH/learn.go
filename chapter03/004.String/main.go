package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	{
		var a rune = 'a'
		var b rune = '\x61'
		var c rune = '\141'
		var d rune = '\u1234'
		var e rune = '\U00012345'
		//var f rune = '\Ufffffffa'
		var s []rune = []rune{a, b, c, d, e}
		fmt.Printf("%+v\n", s)

		//var s []rune = []rune{a, b, c, d, e}
		//fmt.Printf("%+v\n", s) // [97 97 97 4660 74565]
		fmt.Println(string(s))

		name := "abc"
		fmt.Printf("%+v\n", []rune(name)) // [112 105 114 108 111]
		fmt.Printf("%+v\n", string([]rune(name)))
		fmt.Printf("%+v\n", string(name))
	}

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
