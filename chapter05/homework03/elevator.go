package main

import "fmt"

type Elevator struct {
	curFloor     int    // 当前楼层
	curStatus    string // 电梯状态 0: "paused", 1: "stoped", 2: "uping", 3: "downing",
	targetFloors []int
	direction    int
}

func (e *Elevator) getCurFloor() (curFloor int) {
	curFloor = e.curFloor
	return
}

func (e *Elevator) getCurStatus() (curStatus string) {
	curStatus = e.curStatus
	return
}

func (e *Elevator) addTargetFloors(targetFloors ...int) {
	e.targetFloors = append(e.targetFloors, targetFloors[:]...)
}

func (e *Elevator) getTargetFloors() (targetFloors []int) {
	return e.targetFloors
}

func (e *Elevator) startMove() {
	curFloor := e.curFloor
	for _, item := range e.targetFloors {
		space := 0
		if curFloor > item {
			e.direction = elevatorDirectionMap["down"]
			space = curFloor - item
		} else if curFloor < item {
			e.direction = elevatorDirectionMap["up"]
			space = item - curFloor
		}
		fmt.Println(space)
		//time.Sleep((int(space) * time.Second))
	}
	//time.Sleep(2 * time.Second)

	e.curFloor = 3
	// 清空
	//e.targetFloors = append(e.targetFloors, e.targetFloors[len(e.targetFloors)-1:]...)
}

func (e *Elevator) clearTargetFloors() {
	e.targetFloors = append([]int{})
}
