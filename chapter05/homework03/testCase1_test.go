package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	/*
		楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
	*/
	{
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

		curStatus := eleSvc.elevator.getCurStatus()
		fmt.Println(curStatus)
		wantCurStatus := elevatorStatusMap["stopped"]
		if curStatus != wantCurStatus {
			t.Fatalf("预期电梯%v， 但得到的是%v", wantCurStatus, curStatus)
		}
	}
}
