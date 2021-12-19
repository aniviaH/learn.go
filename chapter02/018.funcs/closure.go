package main

import "fmt"

var counter int

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	//
	return
}

func showUseTimes() {
	fmt.Println("counter:", counter)
}

func getImprovementFunc() func(percentage float64) {
	base := 1000.0

	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println("newBase:", base)
	}
}
func closureMain() {
	imp := getImprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)
}
