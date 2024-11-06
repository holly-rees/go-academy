package assignment6

import "testing"

func TestConvertStringToAge(t *testing.T) {

}
func TestCalculateAge(t *testing.T) {
	expected := 32

	actual := CalculateAge()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
