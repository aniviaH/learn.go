package main

import "fmt"

func main() {
	names := []string{"小强", "周一", "李二"}
	fr := []float64{20, 24, 25}

	names = append(names, "张三")
	fr = append(fr, 27)

	for i, name := range names {
		if name == "张三" {
			fmt.Printf("%s 的体脂率是：%f\n", name, fr[i])
		}
	}

	fmt.Println("定义map-----")
	/*
		Map注意事项
		1.Map 不用实例化就可以读取、删除
		2.Map必须 实例化 才可以写入
		3.重复删除相同的Key不会引起异常
	*/
	var m1 map[string]int = nil
	//m1["a"] = 1 // panic: assignment to entry in nil map
	delete(m1, "a")
	fmt.Println("m1 没有实例化，直接读取", m1["a"])

	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "李静": 70, "苗文": 80}
	fmt.Println(m1, m2, m3)

	fmt.Println("王强的分数：", m3["王强"])
	fmt.Println("小强的分数：", m3["小强"])

	// 判断key是否真的存在于Map中
	xqScore, ok := m3["小强"] // 假分数-默认值
	fmt.Println(xqScore, "------", ok)

	m3["王强"] = 100
	fmt.Println(m3)
	delete(m3, "王强")
	fmt.Println(m3)

	m3["小强"] = 90
	fmt.Println(m3)
	fmt.Println("小强的分数：", m3["小强"])
	xqScore, ok = m3["小强"] // 真分数-key存在于Map
	fmt.Println(xqScore, "------", ok)

	for k, v := range m3 {
		fmt.Printf("%s\t%d\n", k, v)
	}

	//mSurprise := map[string]map[int]map[int]map[int]string{}

	var m4 map[int][]string = map[int][]string{11: []string{"苹果", "香蕉"}}
	fmt.Println(m4)
	m5 := map[string]map[int]float64{}
	fmt.Println(m5)

	m6 := map[string]int{"a": 1, "b": 2}

	//type Map = map[int]string
	m7 := map[string]int{}
	m7 = m6
	fmt.Println(m7)

	//var m10 map[string][]int
	//delete(m10, "a")
	//m1["a"] = 1 // 程序异常退出
	//fmt.Println(m1)
	//var m10 map[int][]string
	//m10[2] = []string{"a"} // 程序异常退出
	//fmt.Println(m1)
}
