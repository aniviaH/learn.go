package calculator

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	var err error
	fatRate, err = gobmi.CalcFatRate(bmi, age, sex)
	fmt.Println(err)
	return fatRate
}
