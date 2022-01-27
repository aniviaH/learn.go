package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	c := []int{4, 5}
	fmt.Println(a) // a 是切片
	fmt.Println("追加元素到a中，a是切片")
	//a = append(a, 333, 444, 555)
	a = append(a, c...)
	fmt.Println(a)

	b := [0]int{} // b 是数组
	fmt.Println(b)
	fmt.Println("追加元素到b中，b是数组")
	//b = append(b, 333) // 编译错误，无法追加到固定长度数组

	xqInfo := []string{"小强", "男", "在职"}
	fmt.Println(xqInfo)
	for i, v := range xqInfo {
		fmt.Println(i, v)
	}
	fmt.Println(xqInfo[0])

	fmt.Println("-----------")
	fmt.Println("删除切片中的元素")
	a = []int{111, 222, 333, 444, 555}
	fmt.Println("删除前:：", a)
	a = append(a[:1], a[2:4]...)
	//a = append(a[2:], a[3:]...)
	fmt.Println("删除后：", a)
	a = append(a, a[:]...)
	fmt.Println("double后：", a)

	backup := a[1:]
	//backup := append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	fmt.Println(a, backup) // 因为backup的访问和a是相同访问的引用数据，指向同一片内存地址
	e := append(a, backup...)
	fmt.Println(e)

	//
	fmt.Println("删除索引 2 处的元素-->>")
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	index := 2
	//arr = append(arr[:index], arr[index+1:]...)
	arr = removeSliceItemAtIndex(arr, index)
	fmt.Println(arr)

	// 复制
	arr2 := make([]int, 4)
	copy(arr2, arr[1:])
	fmt.Println(arr2)
}

func removeSliceItemAtIndex(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}
