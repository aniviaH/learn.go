package main

import "fmt"

/*
计算多个人的平均体脂
连续输入多人的体脂计算信息，最后输出所有人的平均体脂
*/
func main() {
	fmt.Println("作业01-1：计算多个人的平均体脂，连续输入多人的体脂计算信息，最后输出所有人的平均体脂")

	const UserCount int = 3
	fmt.Printf("请输入%d位用户信息：\n", UserCount)

	var names = [UserCount]string{}
	var fateRates = [UserCount]float64{}
	var fateRateTotal float64
	var fateRateAverage float64

	for i := 0; i < UserCount; i++ {
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

		var sexRate int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexRate = 1
		} else {
			sexRate = 0
		}

		var age int
		fmt.Print("年龄：")
		fmt.Scanln(&age)

		var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexRate)) / 100

		names[i] = name
		fateRates[i] = fatRate
	}

	for _, val := range fateRates {
		fateRateTotal += val
	}
	fateRateAverage = fateRateTotal / float64(UserCount)

	fmt.Println("用户：", names)
	fmt.Println("体脂：", fateRates)
	fmt.Println("平均：", fateRateAverage)
}
