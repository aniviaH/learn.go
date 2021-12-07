package main

import "fmt"

func main() {
	fmt.Print("亲，请输入需要计算体脂的人数：")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("请连续 %d 次输入次用户信息\n", num)

	var names = [3]string{}
	var weights = [3]float64{}
	var heights = [3]float64{}
	var ages = [3]int{}
	var sexes = [3]string{}
	var fateRates = [num]float64{}

	for i := 0; i < num; i++ {
		fmt.Printf("第%d位用户===>>>\n", i+1)

		var name string
		fmt.Print("姓名：")
		fmt.Scanln(&name)

		var weight float64
		fmt.Print("体重（kg）：")
		fmt.Scanln(&weight)

		var height float64
		fmt.Print("身高（米）：")
		fmt.Scanln(&height)

		var bmi float64 = weight / (height * height)

		var age int
		fmt.Print("年龄：")
		fmt.Scanln(&age)

		var sexRate int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexRate = 1
		} else {
			sexRate = 0
		}
		var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexRate)) / 100

		names[i] = name
		weights[i] = weight
		heights[i] = height
		ages[i] = age
		sexes[i] = sex
		fateRates[i] = fatRate
	}

	fmt.Println("names", names)
	fmt.Println("weights", weights)
	fmt.Println("heights", heights)
	fmt.Println("ages", ages)
	fmt.Println("sexes", sexes)
	fmt.Println(fateRates)
}
