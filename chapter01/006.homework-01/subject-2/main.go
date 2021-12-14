package main

import "fmt"

/*
判断两条线是否平行
提示：
• 两点决定一条直线
• 两条线是否平行取决于两条线的斜率是否一样
*/
func main() {
	fmt.Println("作业01-2：判断两条线是否平行")
	// 1.线a 两点
	// 2.线b 两点
	// 3.判断两者斜率是否一致
	var lineA [2][2]float64
	var lineB [2][2]float64

	var ax1, ay1, ax2, ay2 float64
	var bx1, by1, bx2, by2 float64

	fmt.Print("请输入第一条线的起点坐标x：")
	fmt.Scanln(&ax1)

	fmt.Print("请输入第一条线的起点坐标y：")
	fmt.Scanln(&ay1)

	fmt.Print("请输入第一条线的终点坐标x：")
	fmt.Scanln(&ax2)

	fmt.Print("请输入第一条线的终点坐标y：")
	fmt.Scanln(&ay2)

	fmt.Print("请输入第二条线的起点坐标x: ")
	fmt.Scanln(&bx1)

	fmt.Print("请输入第二条线的起点坐标y：")
	fmt.Scanln(&by1)

	fmt.Print("请输入第二条线的终点坐标x：")
	fmt.Scanln(&bx2)

	fmt.Print("请输入第二条线的终点坐标y：")
	fmt.Scanln(&by2)

	lineA[0][0] = ax1
	lineA[0][1] = ay1
	lineA[1][0] = ax2
	lineA[1][1] = ay2
	lineB[0][0] = bx1
	lineB[0][1] = by1
	lineB[1][0] = bx2
	lineB[1][1] = by2

	var shortAY = ay2 - ay1
	var shortAX = ax2 - ax1

	var shortBY = by2 - by1
	var shortBX = bx2 - bx1

	fmt.Println("第一条线为：", lineA)
	fmt.Println("第二条线为：", lineB)

	if shortAX == 0 && shortAY == 0 {
		fmt.Println("第一条线的起点终点为同一点（第一条线是一个点）")
		return
	}
	if shortBX == 0 && shortBY == 0 {
		fmt.Println("第二条线的起点终点为同一点（第二条线是一个点）")
		return
	}

	if shortAX == 0 && shortBX == 0 {
		fmt.Println("平行，两条线是竖直线")
	}

	var slopeA = shortAY / shortAX
	var slopeB = shortBY / shortBX

	fmt.Println("第一条线的斜率：", slopeA)
	fmt.Println("第二条线的斜率：", slopeB)
	if slopeA-slopeB == 0 {
		fmt.Println("两条线平行")
	} else {
		fmt.Println("两条线不平行")
	}
}
