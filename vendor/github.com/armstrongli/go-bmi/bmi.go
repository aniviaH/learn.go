package gobmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	fmt.Println("staging---g0-bmi-----v-2--BMI")
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}
