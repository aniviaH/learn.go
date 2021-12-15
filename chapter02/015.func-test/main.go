package main

import (
	"fmt"
)

func main() {
	//starttime := time.Now()
	//defer func() {
	//	fmt.Println("defer end:", time.Now().Sub(starttime))
	//}()
	//
	//time.Sleep(5 * time.Second)
	//fmt.Println("end", time.Now())

	c, d, e, f := minus(1, 2)
	fmt.Println(c, d, e, f)
}

func minus(a, b int) (c, d, e, f int) {
	return a - b, c + d, -1, f + e
}
