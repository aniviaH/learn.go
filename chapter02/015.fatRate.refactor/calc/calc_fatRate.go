package calc

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
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
