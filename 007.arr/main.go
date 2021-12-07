package main

import "fmt"

func main() {
	var age int = 25
	fmt.Println(age)

	var ages1 = [5]int{25, 35, 45, 55}
	ages2 := [5]int{25, 35, 45, 55}
	fmt.Println(ages1)
	fmt.Println(ages2)

	ages3 := [6]int{1, 2, 3, 4, 5}
	fmt.Println(ages3)
	//ages2 = ages3

	ages2 = [...]int{1, 2, 3, 4, 5}

	var ages4 = [4]int{}
	fmt.Println("ages4", ages4)

	var a [3]int = [3]int{}
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("a =", a, " len = ", len(a))
	//a[-1] = 0 // 编译错误，越界
	//a[4] = 4 // 编译错误，越界

	fmt.Println("遍历======for")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("遍历======for range")
	for i := range a {
		fmt.Println(a[i])
	}
	for i, val := range a {
		//fmt.Println("val = ", val)
		fmt.Println("i ===>>> ", i, " val ===>>> ", val)
	}
	for _, val := range a {
		// 不想用第一个下标值时
		fmt.Println("val = ", val)
	}
}
