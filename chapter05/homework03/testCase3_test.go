package main

import (
	"fmt"
	"testing"
)

func TestCase3(t *testing.T) {
	/*
		楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
	*/
	{
		// 初始化
		eleSvc := &ElevatorService{
			building: &Building{
				floors:              []int{0, 1, 2, 3, 4},
				floorsHasPressedBtn: []int{},
			},
			elevator: &Elevator{
				curFloor:     2,
				curStatus:    elevatorStatusMap["stopped"],
				direction:    elevatorDirectionMap["stopped"],
				targetFloors: []int{},
			},
		}
		fmt.Println(eleSvc.building)
		fmt.Println(eleSvc.elevator)

		// 操作开始
		floorOfPressBtn1 := 4
		indexFloorOfPressBtn1 := floorOfPressBtn1 - 1
		eleSvc.building.pressBtn(indexFloorOfPressBtn1)
		floorOfPressBtn2 := 2
		indexFloorOfPressBtn2 := floorOfPressBtn2 - 1
		eleSvc.building.pressBtn(indexFloorOfPressBtn2)

		eleSvc.elevator.start(eleSvc.building)

		curFloor := eleSvc.elevator.getCurFloor()
		wantCurFloorIndex := 2 - 1
		if curFloor != wantCurFloorIndex {
			t.Fatalf("预期电梯停在%v楼， 但停在%v楼", wantCurFloorIndex+1, curFloor+1)
		}
	}
}
