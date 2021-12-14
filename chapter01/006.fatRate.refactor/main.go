package main

import "fmt"

func main() {
	/*
		1.录入一些******
		name, weight, height, age, sex, ... = getFromInput()
		2.计算出体脂率
		fateRate = calculateFateRate(name, weight, height, ...)
		3.给出体脂健康与建议
		suggestion = checkHealthStateAndGetSuggestion(fateRate, sex, age)
	*/
	for {
		_, weight, height, age, sex := getFromInput()

		bmi := calcBMI(weight, height)

		var maleFatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(1)) / 100
		var femaleFatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(0)) / 100
		var fatRate float64
		if sex == "男" {
			fatRate = maleFatRate
		} else {
			fatRate = femaleFatRate
		}

		fmt.Println("体脂率是：", fatRate)

		if fatRate < 0 {
			fmt.Println("输入项有误，请检查重新输入！")
			continue
		}

		checkHealthStateAndGetSuggestion(sex, age, fatRate)

		var whetherContinue string
		fmt.Print("是否录入下一个(y/n)? ")
		fmt.Scanln(&whetherContinue)
		if whetherContinue == "y" {
			//continue
		} else {
			break
		}
	}
}

func checkHealthStateAndGetSuggestion(sex string, age int, fatRate float64) {
	if sex == "男" {
		// todo 编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else if fatRate > 0.27 {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		} else if age >= 60 {
			if fatRate <= 0.13 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else if fatRate > 0.29 {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
		}
	} else {
		// todo 编写女性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		} else if age >= 60 {
			if fatRate <= 0.22 {
				fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				fmt.Println("目前是：标准。太棒了，要保持哦！")
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				fmt.Printf("目前是：偏重。吃完饭多散散步，消化消化。")
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				fmt.Println("目前是：肥胖。少吃点，多运动运动。")
			} else {
				fmt.Println("目前是：严重肥胖。健身游泳跑步，即刻开始。")
			}
		}
	}
}

func calcBMI(weight float64, height float64) float64 {
	var bmi float64 = weight / (height * height)
	return bmi
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
