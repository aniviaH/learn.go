package main

import "fmt"

func main() {
	sameSubdomain()
	sameSubdomain2()

	sampleSubdomainInIf()
	sampleSubdomainInFor()
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
