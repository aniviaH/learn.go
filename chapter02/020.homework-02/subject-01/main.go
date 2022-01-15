package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {
	/*
		录入一个人的各项指标
		计算体脂
		给出建议
	*/
	for {
		//testRes := test()
		calcFatRateBody()

		// 判断继续
		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func calcFatRateBody() {
	// 录入各项
	weight, height, age, _, sex := getMaterialsFromInput()

	// 计算体脂率
	//fatRate := calcFatRate(weight, height, age, sexRate, sex)

	bmi, err := gobmi.BMI(weight, height)
	if err != nil {
		fmt.Println("计算BMI异常", err)
		return
	}
	fatRate, _ := gobmi.CalcFatRate(bmi, age, sex)
	fmt.Println("体脂率是：", fatRate)

	if fatRate < 0 {
		fmt.Println("输入项有误，请检查重新输入！", err)
		panic("fat rate is not allowed to be negative")
	}

	// 得到健康提示
	var checkHealthFunc func(age int, fatRate float64)
	if sex == "男" {
		// todo 编写男性的体脂率与体脂状态表
		// fixme xxx-xxx
		checkHealthFunc = getHealthSuggestionForMale
	} else {
		// todo 编写女性的体脂率与体脂状态表
		checkHealthFunc = getHealthSuggestionForFemale
	}
	checkHealthFunc(age, fatRate)
	return
}

func getHealthSuggestionForFemale(age int, fatRate float64) {
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

func getHealthSuggestionForMale(age int, fatRate float64) {
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
}

func getMaterialsFromInput() (float64, float64, int, int, string) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（kg）：")
	fmt.Scanln(&weight)

	var height float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&height)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	var sexRate int
	if sex == "男" {
		sexRate = 1
	} else {
		sexRate = 0
	}

	return weight, height, age, sexRate, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个(y/n)? ")
	fmt.Scanln(&whetherContinue)
	if whetherContinue == "y" {
		//continue
		return true
	} else {
		return false
	}
}
