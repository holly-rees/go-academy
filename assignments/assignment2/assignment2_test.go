package assignments

import (
	"testing"
)

func TestConcatenateNames(t *testing.T) {
	expected := "Holly Clare Rees"
	actual := ConcatenateNames("Holly", "Clare", "Rees")
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
