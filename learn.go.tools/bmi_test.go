package gobmi

import "testing"

func TestBMI1(t *testing.T) {
	_, err := BMI(-1, 1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestBMI2(t *testing.T) {
	_, err := BMI(1, -1)
	if err == nil {
		t.Errorf("should be error, but err returned is nil")
	}
}

func TestBMI3(t *testing.T) {
	var expectedBmp float64 = 1
	bmi, err := BMI(1, 1)
	if err != nil {
		t.Errorf("should be error, but err returned is nil")
	}
	if bmi != expectedBmp {
		t.Error("expected is different from calculated")
	}
}
