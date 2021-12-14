package main

import "fmt"

func main() {
	{
		fmt.Println("测试代码块1------")
		a := make([]int, 0)
		fmt.Println("len:", len(a), "cap:", cap(a), "val:", a)
		a = append(a, 1)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
	}

	{
		fmt.Println("测试代码块2------")
		a := make([]int, 1)
		fmt.Println("len:", len(a), "cap: ", cap(a), "val:", a)
		a = append(a, 1)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
	}

	{
		fmt.Println("测试代码块3------")
		a := make([]int, 1, 2)
		fmt.Println("len:", len(a), "cap:", cap(a), "val:", a)
		a = append(a, 1)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
		a = append(a, 2, 3)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
		a = append(a, 4)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
		a = append(a, 5, 6, 7)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
		a = append(a, 9)
		fmt.Println("len:", len(a), "cap:", cap(a), "val: ", a)
	}
}
