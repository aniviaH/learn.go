package main

type Building struct {
	floors     []int
	pressedArr []int
}

func (b *Building) pressBtn(floors ...int) (err error) {
	//floorNum := len(b.floors)
	//index := floor - 1
	//if index < 0 || index >= floorNum {
	//	//log.Printf("所按楼层不能小于%v，不能大于%v", 1, floorNum)
	//	err = fmt.Errorf("所按楼层不能小于%v，不能大于%v", 1, floorNum)
	//	return err
	//}
	//b.pressedArr[index] = true
	b.pressedArr = append(b.pressedArr, floors[:]...)
	return nil
}
func (b *Building) clearPressedArr() {
	b.pressedArr = append([]int{})
}
