package assignment6

import "testing"

func TestConvertStringToAge(t *testing.T) {

}
func TestCalculateAge(t *testing.T) {
	expected := 32

	actual := CalculateAge("13/09/1992")

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
