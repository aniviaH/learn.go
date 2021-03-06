package main

import (
	"fmt"
	"time"
)

type Elevator struct {
	curFloor     int    // 当前楼层
	direction    string //电梯行进方向
	curStatus    string // 电梯状态启动还是停止
	targetFloors []int
}

func (e *Elevator) start(b *Building) {
	curStatus := e.curStatus
	if len(b.floorsHasPressedBtn) > 0 {
		if curStatus == elevatorStatusMap["stopped"] {
			// 电梯原本停止
			e.targetFloors = append(e.targetFloors, b.floorsHasPressedBtn...)
			e.move()
		} else if curStatus == elevatorStatusMap["running"] {
			// 电梯原本运行中
			e.targetFloors = append(e.targetFloors, b.floorsHasPressedBtn...)
		}
	} else {
		// 无需启动行进
	}
}

func (e *Elevator) move() {
	length := len(e.targetFloors)
	if length == 0 {
		e.curStatus = elevatorStatusMap["stopped"]
		e.direction = elevatorDirectionMap["stopped"]
		return
	}

	firstTarget := e.targetFloors[0]
	curFloor := e.curFloor
	step := 0
	if curFloor > firstTarget {
		e.direction = elevatorDirectionMap["down"]
		step = curFloor - firstTarget

		e.moveToTargetFloor(firstTarget, step)
	} else if curFloor < firstTarget {
		e.direction = elevatorDirectionMap["up"]
		step = firstTarget - curFloor
		e.moveToTargetFloor(firstTarget, step)
	} else {
		e.direction = elevatorDirectionMap["stopped"]
	}
}

func (e *Elevator) moveToTargetFloor(target int, step int) {
	d := e.direction
	var moveOneFloorFunc func(int)
	if d == elevatorDirectionMap["up"] {
		moveOneFloorFunc = e.upOneFloor
	} else if d == elevatorDirectionMap["down"] {
		moveOneFloorFunc = e.downOneFloor
	}

	fmt.Printf("电梯在%v楼，向%v楼行进，方向：%v\n", e.curFloor+1, target+1, elevatorDirectionTextMap[d])
	for i := 0; i < step; i++ {
		moveOneFloorFunc(target)
	}
}

func (e *Elevator) upOneFloor(target int) {
	fmt.Println("上升一楼")

	// todo 这里可以中途按钮的场景

	time.Sleep(1 * time.Second)

	e.curFloor++

	e.openDoor()
	e.closeDoor()

	fmt.Println("当前楼层：", e.curFloor+1)

	e.checkIsArrivedOneTargetFloor(target)
	//e.move()
}
func (e *Elevator) downOneFloor(target int) {
	fmt.Println("下降一楼")
	time.Sleep(1 * time.Second)

	e.curFloor--

	e.openDoor()
	e.closeDoor()

	e.checkIsArrivedOneTargetFloor(target)
	//e.move()
}

func (e *Elevator) checkIsArrivedOneTargetFloor(target int) {
	if e.curFloor == target {
		// 到达一个目标楼层，移除第一个目标层
		e.targetFloors = append([]int{}, e.targetFloors[1:]...)

		e.move()
	}

	// todo 中途按的按钮，放在数组后面，到达楼层后需要需要进行判断，将其移除
}

func (e *Elevator) openDoor() {
	fmt.Println("打开电梯门")
}
func (e *Elevator) closeDoor() {
	fmt.Println("关闭电梯门")
}

func (e Elevator) stop() {

}

func (e Elevator) pauseInOneFloor() {

}

// get
func (e Elevator) getCurFloor() int {
	return e.curFloor
}
func (e Elevator) getCurStatus() string {
	return e.curStatus
}
