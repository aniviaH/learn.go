package main

import (
	"fmt"
)

func main() {
	fmt.Printf("我的名字是%s\n", "小明")
	fmt.Printf("我的名字是%q\n", "小明")
	fmt.Printf("我的名字是%x\n", "小明")
	fmt.Printf("我的名字是%X\n", "小明")

	// 课后作业：计算两个圆的面积之和并输出，精确到小数点后3 位。
	var width float64 = 2.1
	var height float64 = 4
	var rect = width * height
	fmt.Printf("方形面积：%b\n", rect)

	fmt.Printf("方形面积：%e\n", rect)
	fmt.Printf("方形面积：%E\n", rect)

	fmt.Printf("方形面积：%.3f\n", rect)
	fmt.Printf("方形面积：%.3F\n", rect)

	fmt.Printf("方形面积：%g\n", rect)
	fmt.Printf("方形面积：%G\n", rect)
}
