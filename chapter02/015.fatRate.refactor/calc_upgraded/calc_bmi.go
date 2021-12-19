package calc

import "fmt"

func init() {
	fmt.Println("init---函数可以做包的初始化。比如：全局变量的初始化、注册、预先计算等。")
}

func CalcBMI(height float64, weight float64) (bmi float64) {
	if height <= 0 {
		panic("身高不能是0或负数")
	}
	if weight <= 0 {
		panic("体重不能是0或负数")
	}
	return weight / (height * height)
}
