package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {
	//fmt.newPrinter()

	panicAndRecover()

	deferGuess()

	openFile()
	fmt.Println("slee somewhile--")
	time.Sleep(10 * time.Second)

	closureMain()
	fmt.Println("slee somewhile--")
	time.Sleep(10 * time.Second)

	fmt.Println(calcSum(1, 2, 3))
	fmt.Println(calcSum(1, 2, 3))
	fmt.Println(calcSum(1, 2, 3))
	fmt.Println(calcSum(1, 2, 3))
	fmt.Println(calcSum(1, 2, 3))

	showUseTimes()

	//num := fibonacci(50) // 复杂度：O(2^n)
	//fmt.Println(num)
	guessNumber(1, 100)

	sameSubdomain()
	sameSubdomain2()

	sampleSubdomainInIf()
	sampleSubdomainInFor()

	fmt.Println("全局变量赋值前：")
	calcAdd() // 0
	tall, weight = 1.7, 70
	fmt.Println("全局变量赋值后：")
	calcAdd() // 71.7

	// 重新定义重名的局部变量
	tall, weight := 100, 70.0
	calcAdd() // 71.7
	tall, weight = 200, 70.0
	calcAdd() // 71.7
	fmt.Println(tall, weight)
}

func sameSubdomain() {
	name := "张三"
	fmt.Println("名字---0：", name) // 张三
	{
		fmt.Println("{}内名字---1：", name) // 张三
		name = "李四"
		fmt.Println("{}内名字---1：", name) // 李四
	}
	fmt.Println("{}外名字---3", name) // 李四
}

func sameSubdomain2() {
	name := "张三"
	fmt.Println("名字---4：", name) // 张三
	{
		fmt.Println("{}内名字---5：", name) // 张三
		name := "李四"
		fmt.Println("{}内名字---6：", name) // 李四
	}
	fmt.Println("{}外名字---7", name) // 张三
}

func sampleSubdomainInIf() {
	if v := calcBMI(); v == 0 {
		fmt.Println("if---", v)
	} else {
		fmt.Println("else---", v)
	}
	//fmt.Println(v) //无效 v的有效范围为 if block
}
func sampleSubdomainInFor() {
	for i := 0; i < 3; i++ {
		fmt.Println("for内部---", i)
	}
	//fmt.Println("for外部---", i) // i的有效范围为 for block
}

func calcBMI() (bmi int) {
	//var bmi float64 = weight / (height * height)
	fmt.Println("calcBMI")
	return 1
}

func calcAdd() float64 {
	fmt.Println(tall + weight)
	return 1
}
