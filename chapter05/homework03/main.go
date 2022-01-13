package main

import "fmt"

var (
	elevatorStatusMap = map[int]string{
		//0: "stoped",
		//1: "paused",
		//2: "uping",
		//3: "downing",
		0: "停止",
		1: "暂停",
		2: "上升中",
		3: "下降中",
	}
)

func main() {
	// 初始化
	building := &Building{
		floors: []int{1, 2, 3, 4, 5},
	}
	elevator := &Elevator{
		curFloor:  1,
		curStatus: elevatorStatusMap[1],
	}

	fmt.Println(*building)
	fmt.Println(*elevator)

	fmt.Println("电梯当前楼层：", elevator.getCurFloor())
	// 获取初始化之后的状态
}
