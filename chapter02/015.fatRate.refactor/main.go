package main

import (
	"fmt"
	c02 "learn.go/chapter02/015.fatRate.refactor/calc_upgraded"
	_ "testing/quick"
	//_ "learn.go/chapter02/001.fatRate.refactor/calc" // 不需要被调用，但它是项目的需要
	c01 "learn.go/chapter02/015.fatRate.refactor/calculator" // 引用包并起别名
	//. "learn.go/chapter02/001.fatRate.refactor/calc" // 扩展当前包
)

/*
init 函数
• init 函数是在包被引用时用于包初始化的函数。
特殊点：
• 不需要，也不可以被调用。Golang 默认自动执行。
• 一个Go 文件、包中可以有多个init 函数。
• 在main函数之前执行
*/
func init() {
	fmt.Println("init函数---1")
}
func init() {
	fmt.Println("init函数---2")
}
func init() {
	fmt.Println("init函数---3")
}

func main() {
	fmt.Print("main---")
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

func test() int {
	return 1
}

func calcFatRateBody() {
	// 录入各项
	weight, height, age, sexRate, sex := getMaterialsFromInput()

	// 计算体脂率
	fatRate := calcFatRate(weight, height, age, sexRate, sex)
	//bmi := calc.CalcBMI(height, weight)
	//fatRate := calc.CalcFatRate(bmi, age, sex)
	//bmi := calcAlia.CalcBMI(height, weight)
	//fatRate := calcAlia.CalcFatRate(bmi, age, sex)

	if fatRate < 0 {
		fmt.Println("输入项有误，请检查重新输入！")
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
	//getHealthSuggestions(checkHealthFunc)
	//getHealthSuggestions(age, fatRate, getHealthSuggestionForMale)
	//getHealthSuggestions(age, fatRate, getHealthSuggestionForFemale)
	//getHealthSuggestions(age, fatRate, generateCheckHealthSuggestionForMale)
}

func getHealthSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}

func generateCheckHealthSuggestionForMale() func(age int, fatRate float64) {
	// 定制
	return func(age int, fatRate float64) {

	}
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

func calcFatRate(weight float64, height float64, age int, sexRate int, sex string) float64 {
	//bmi := weight / (height * height)
	bmi := c01.CalcBMI(height, weight)
	//fatRate := (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexRate)) / 100
	fatRate := c02.CalcFatRate(bmi, age, sex)
	fmt.Println("体脂率是：", fatRate)
	return fatRate
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
