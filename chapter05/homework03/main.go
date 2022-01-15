package main

import "fmt"

var (
	elevatorStatusMap = map[string]string{
		"stopped": "stopped",
		"running": "running",
	}

	elevatorDirectionMap = map[string]string{
		"stopped": "stopped",
		"up":      "up",
		"down":    "down",
	}
	elevatorDirectionTextMap = map[string]string{
		"stopped": "停止",
		"up":      "上升",
		"down":    "下降",
	}
)

func main() {
	// 初始化
	eleSvc := &ElevatorService{
		building: &Building{
			floors:              []int{0, 1, 2, 3, 4},
			floorsHasPressedBtn: []int{},
		},
		elevator: &Elevator{
			curFloor:     0,
			curStatus:    elevatorStatusMap["stopped"],
			direction:    elevatorDirectionMap["stopped"],
			targetFloors: []int{},
		},
	}
	fmt.Println(eleSvc.building)
	fmt.Println(eleSvc.elevator)

	fmt.Println("电梯当前楼层：", eleSvc.elevator.getCurFloor())
}
