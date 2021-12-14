package main

import "fmt"

func main() {
	//var m1 = map[string]int{"a": 1, "b": 2}
	//var m2 = map[string]int{"c": 3, "d": 4}
	//
	//for k, v := range m2 {
	//	m1[k] = v
	//}
	//fmt.Println(m1)

	leftMap, rightMap := map[string]int{}, map[string]int{}

	leftMap["语文"] = 80
	rightMap["数学"] = 50

	for k, val := range rightMap {
		leftMap[k] = val
	}
	fmt.Println(leftMap)
}
