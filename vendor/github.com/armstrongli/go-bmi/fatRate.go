package gobmi

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	fmt.Println("CalcFatRate---本地replace+扩展---2")

	if sex != "男" && sex != "女" {
		err = fmt.Errorf("录入性别请输入男/女")
	}

	if age < 0 {
		err = fmt.Errorf("age cannot be negative")
		return
	}
	if age == 0 {
		err = fmt.Errorf("age cannot be 0")
		return
	}
	if age > 150 {
		err = fmt.Errorf("age cannot lager than 150")
		return
	}

	var sexRate int
	if sex == "男" {
		sexRate = 1
	} else {
		sexRate = 0
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexRate)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	return
}
