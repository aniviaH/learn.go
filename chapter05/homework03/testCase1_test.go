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
		building := &Building{
			floors: []int{1, 2, 3, 4, 5},
		}
		fmt.Println("楼层：", building)

		elevator := &Elevator{
			curFloor:  1,
			curStatus: elevatorStatusMap[0],
		}

		curStatus := elevator.getCurStatus()
		fmt.Println(curStatus)
		if curStatus != "停止" {
			t.Fatalf("预期电梯不动， 但得到的是%v", curStatus)
		}
	}
}
