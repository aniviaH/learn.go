package main

import (
	"fmt"
	"log"
	"testing"
)

func TestCase2(t *testing.T) {
	/*
		楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
	*/

	// 初始化
	building := &Building{
		floors:     []int{1, 2, 3, 4, 5},
		pressedArr: []int{},
	}
	elevator := &Elevator{
		curFloor:  1,
		curStatus: elevatorStatusMap[0],
	}

	// 按钮操作
	pressesFloors := []int{3}
	err := building.pressBtn(pressesFloors...)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("楼层：", building)
	// 电梯移动
	elevator.startMove()

	// 清空状态
	elevator.clearTargetFloors()
	building.clearPressedArr()

	curFloor := elevator.getCurFloor()
	curTargetFloors := elevator.getTargetFloors()
	if curFloor != 3 {
		t.Fatalf("预期电梯停在3楼， 但停留在%v楼", curFloor)
	}
	if len(curTargetFloors) != 0 {
		t.Fatalf("预期电梯没有目标楼层， 但还有目标楼层：%v", curTargetFloors)
	}
}
