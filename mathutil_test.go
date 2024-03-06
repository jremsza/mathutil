package mathutil

import (
	"testing"
)

func TestAverage(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 3
	if result := Average(numbers); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestAverageFloat(t *testing.T) {
	numbers := []float64{94.0, 29.0, 6.0, 36.0, 1.0, 48.0, 48.0, 63.0, 81.0, 86.0, 13.0, 45.0, 40.0, 62.0, 71.0, 29.0, 28.0, 68.0, 100.0, 23.0}
	expected := 48.333333333333336
	if result := Average(numbers); result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestAverageInt32(t *testing.T) {
	numbers := []int32{94, 29, 6, 36, 1, 48, 48, 63, 81, 86, 13, 45, 40, 62, 71, 29, 28, 68, 100, 23}
	expected := int32(48)
	if result := Average(numbers); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
