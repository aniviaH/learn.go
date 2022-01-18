package main

import (
	"fmt"
	"testing"
)

func TestCase2(t *testing.T) {
	/*
		楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
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

		// 操作开始
		floorOfPressBtn := 3
		indexFloorOfPressBtn := floorOfPressBtn - 1
		eleSvc.building.pressBtn(indexFloorOfPressBtn)
		eleSvc.elevator.start(eleSvc.building)

		curFloor := eleSvc.elevator.getCurFloor()
		wantCurFloor := 3 - 1
		if curFloor != wantCurFloor {
			t.Fatalf("预期电梯停在%v楼， 但停在%v楼", wantCurFloor+1, curFloor+1)
		}
	}
}
