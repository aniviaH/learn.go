package main

import "fmt"

func main() {
	for {
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

		fmt.Println("体脂率是：", fatRate)

		if fatRate < 0 {
			fmt.Println("输入项有误，请检查重新输入！")
			continue
		}

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
