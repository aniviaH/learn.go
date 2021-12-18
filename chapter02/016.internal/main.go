package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	test()
}

func test() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println("执行耗时：", finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()

	testPanic()

	//arr2 := []int{}
	//arr2 := make([]int, 3)
	arr2 := make([]int, 4, 9)
	fmt.Println("len:", len(arr2), "cap:", cap(arr2))
	//fmt.Println(arr[0])
	fmt.Println("default:", arr2[0])
	fmt.Println("default:", arr2[1])
	fmt.Println("default:", arr2[2])
	//arr2[3] = 3
	arr2 = append(arr2, 3)
	fmt.Println("len:", len(arr2), "cap:", cap(arr2))

	arr3 := make([]int, 4, 9)
	//for i := 0; i < len(arr2); i++ {
	//	arr3[i] = arr2[i]
	//}
	copy(arr3, arr2)
	fmt.Println("arr3: ", arr3)

	var pint = new(int)
	var pstring = new(string)
	var pintVal **int = &pint
	var pstringVal = &pstring
	fmt.Println(reflect.TypeOf(pint))
	fmt.Println(reflect.TypeOf(pstring))
	fmt.Printf("pint: %p\n", pint)
	fmt.Printf("pstring: %p\n", pstring)
	fmt.Printf("pintVal: %p\n", pintVal)
	fmt.Printf("pstringVal: %p\n", pstringVal)

	var i = 333
	var j = &i
	fmt.Println("j类型：", reflect.TypeOf(j))
	fmt.Printf("j指针%p:\n", j)

}

func testPanic() {
	var arr3 = make([]int, 0, 4)
	arr3[0] = 0
}
