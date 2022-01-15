package calculator

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBMI(height float64, weight float64) (bmi float64) {
	var err error
	bmi, err = gobmi.BMI(weight, height)
	fmt.Println(err)
	return bmi
}
