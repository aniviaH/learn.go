package main

import (
	"fmt"
	"testing"
)

func TestCase3(t *testing.T) {
	/*
		楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
	*/

	building := &Building{
		floors: []int{1, 2, 3, 4, 5},
	}
	fmt.Println("楼层：", building)

	elevator := &Elevator{
		curFloor:  2,
		curStatus: elevatorStatusMap[0],
	}

	newTargetFloors := []int{4, 2}
	elevator.addTargetFloors(newTargetFloors...)
	elevator.startMove()

	curFloor := elevator.getCurFloor()
	curTargetFloors := elevator.getTargetFloors()
	if curFloor != 2 {
		t.Fatalf("预期电梯停在2楼， 但停留在%v楼", curFloor)
	}
	if len(curTargetFloors) != 0 {
		t.Fatalf("预期电梯没有目标楼层， 但还有目标楼层：%v", curTargetFloors)
	}
}
