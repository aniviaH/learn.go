package main

type Elevator struct {
	curFloor  int    // 当前楼层
	curStatus string // 电梯状态 0: "paused", 1: "stoped", 2: "uping", 3: "downing",
}

func (e *Elevator) getCurFloor() (curFloor int) {
	curFloor = e.curFloor
	return
}

func (e *Elevator) getCurStatus() (curStatus string) {
	curStatus = e.curStatus
	return
}
