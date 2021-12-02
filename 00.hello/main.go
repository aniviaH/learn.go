package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go!!!")

	var name = "刘欢"

	{
		val, err := strconv.Atoi(name)
		fmt.Println(val, err)
	}

	age := 3

	fmt.Printf("%p\n", &age)

	//age := 3
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)

	a := 1

	a, b := 2, 3

	fmt.Println(name, age, time, a, b)

	{
		age := 3
		fmt.Printf("%p\n", &age)
	}

	sha256.New()
}
