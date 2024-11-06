package assignments

import (
	"testing"
)

func TestIsBetween1and10(t *testing.T) {
	expected := true
	actual := isBetween1and10(5)
	if actual != expected {
		t.Errorf("Expected %t but got %t", expected, actual)
	}
}
