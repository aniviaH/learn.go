package main

type Building struct {
	floors              []int
	floorsHasPressedBtn []int
}

func (b *Building) pressBtn(index int) {
	b.floorsHasPressedBtn = append(b.floorsHasPressedBtn, index)
}
