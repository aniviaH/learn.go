package main

import "testing"

func TestWithDirection(t *testing.T) {
	c := make(chan int, 1)
	//iCh := make(chan<- int, 1)
	//oCh := make(<-chan int, 1)
	inOnly(c)
	outOnly(c)
	//inOnly(iCh)
	//inOnly(oCh)
	//outOnly(iCh)
	//outOnly(oCh)

}

func inOnly(c chan<- int) {
	c <- 1
	//<-c // 当c是单向(入)channel时，不能再取数。编译错误
}

func outOnly(c <-chan int) {
	//c <-1 // 当c是单向(出)channel时，不能再存入数据。编译错误
	<-c
}
