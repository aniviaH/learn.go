package main

import "fmt"

func main() {
	constructHellos("张三", "李四")

	// 参数为定长时，必须传参
	avgBmi := calcAvgOnSlice([]float64{1, 2, 3})
	fmt.Println(avgBmi)

	// 不定长参数，参数传递可以更灵活
	bmis := []float64{1.2, 2.1, 1.4, 1.3}
	avgBmi1 := calcAvgBmi(bmis...)
	avgBmi2 := calcAvgBmi(1, 2, 3, 4, 5)
	avgBmi3 := calcAvgBmi([]float64{}...)
	avgBmi4 := calcAvgBmi()
	fmt.Println(avgBmi1)
	fmt.Println(avgBmi2)
	fmt.Println(avgBmi3)
	fmt.Println(avgBmi4)

	// 命名返回值
	getScoreOfStudent("张三")
}

// 不定长参数
func calcAvgBmi(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis))
	return
}

func calcAvgOnSlice(bmis []float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func constructHellos(name ...string) string {
	//fmt.Printf("hello, %s", name)
	return fmt.Sprintf("hello, %s", name)
}

func getScoreOfStudent(stdName string) (chinese int, math int, english int, physics int, nature int) {
	chinese, math, english, physics, nature = 0, 0, 0, 0, 0
	//return chinese, math, english, physics, nature
	return
}
