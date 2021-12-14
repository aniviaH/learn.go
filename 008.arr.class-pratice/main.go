package main

import "fmt"

func main() {
	/*
		课后作业
		• 创建一个一维数组，反转它

		• 用多维数组写一个日历表
		需要考虑每个月的天数不同
		需要考虑是平年还是闰年来专门处理二月
		• 【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列
	*/
	arr := []int{1, 2, 3, 4, 9, 5, 6, 7}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	reverseArr(arr)
	reverseArr(arr2)

}

func reverseArr(arr []int) {
	fmt.Println("反转前：", arr)
	length := len(arr)
	mid := length / 2
	p := 0
	//for {
	//	if p >= mid {
	//		break
	//	}
	//
	//	tmp := arr[p]
	//	arr[p] = arr[length-p-1]
	//	arr[length-p-1] = tmp
	//
	//	p++
	//}

	for ; p < mid; p++ {
		tmp := arr[p]
		arr[p] = arr[length-p-1]
		arr[length-p-1] = tmp
	}

	fmt.Println("反转后：", arr)
}
