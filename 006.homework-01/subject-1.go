package main

import "fmt"

/*
计算多个人的平均体脂
连续输入多人的体脂计算信息，最后输出所有人的平均体脂
*/
func main() {
	fmt.Println("作业01-1：计算多个人的平均体脂，连续输入多人的体脂计算信息，最后输出所有人的平均体脂")

	//const UserCount int = 3
	const UserCount int = 1
	fmt.Printf("请输入%d位用户信息：\n", UserCount)

	var names = [UserCount]string{}
	var weights = [UserCount]float64{}
	var heights = [UserCount]float64{}
	var sexes = [UserCount]string{}
	var ages = [UserCount]int{}

	var fateRates = [UserCount]float64{}
	var suggestions = [UserCount]string{}
	var fateRateTotal float64
	var fateRateAverage float64

	for i := 0; i < UserCount; i++ {
		fmt.Printf("第%d位用户===>>>\n", i+1)

		name, weight, height, age, sex := getFromInput()

		bmi := calcBMI(height, weight)

		var fatRate float64
		var maleFatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(1)) / 100
		var femaleFatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(0)) / 100

		if sex == "男" {
			fatRate = maleFatRate
		} else {
			fatRate = femaleFatRate
		}

		var suggestion string

		if sex == "男" {
			// todo 编写男性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else if fatRate > 0.27 {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else if age >= 60 {
				if fatRate <= 0.13 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.13 && fatRate <= 0.19 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.19 && fatRate <= 0.24 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.24 && fatRate <= 0.29 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else if fatRate > 0.29 {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else {
				suggestion = "不支持评判"
				//fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
			}
		} else {
			// todo 编写女性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.21 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else if age >= 60 {
				if fatRate <= 0.22 {
					suggestion = "偏瘦"
					//fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.22 && fatRate <= 0.29 {
					suggestion = "标准"
					//fmt.Println("目前是：标准。太棒了，要保持哦！")
				} else if fatRate > 0.29 && fatRate <= 0.36 {
					suggestion = "偏重"
					//fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
				} else if fatRate > 0.36 && fatRate <= 0.41 {
					suggestion = "肥胖"
					//fmt.Println("目前是：肥胖。少吃点，多运动运动。")
				} else {
					suggestion = "严重肥胖"
					//fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
				}
			} else {
				suggestion = "不支持评判"
				//fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
			}
		}

		names[i] = name
		weights[i] = weight
		heights[i] = height
		sexes[i] = sex
		ages[i] = age

		fateRates[i] = fatRate
		suggestions[i] = suggestion
	}

	for i := 0; i < UserCount; i++ {
		fmt.Println("用户", i+1, "===>>> 姓名:", names[i], ", 性别:", sexes[i], ", 年龄:", ages[i], ", 身高:", heights[i], ", 体重:", weights[i], ", 体脂:", fateRates[i], ", 建议:", suggestions[i])
	}

	for _, val := range fateRates {
		fateRateTotal += val
	}
	fateRateAverage = fateRateTotal / float64(UserCount)

	//fmt.Println("用户：", names)
	//fmt.Println("体脂：", fateRates)
	fmt.Println(UserCount, "位用户平均体脂：", fateRateAverage)
}

func getFromInput() (name string, weight float64, height float64, age int, sex string) {
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	fmt.Print("身高（米）：")
	fmt.Scanln(&height)

	fmt.Print("年龄：")
	fmt.Scanln(&age)

	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return
}

func calcBMI(height float64, weight float64) float64 {
	return weight / (height * height)
}
