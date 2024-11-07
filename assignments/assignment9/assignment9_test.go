package assignment9

import (
	h "go-assignments/assignments"
	"testing"
)

func TestDisplayFullName(t *testing.T) {
	expected := "First name: Holly\nMiddle name: Clare\nLast name: Rees\n"
	actual := h.CaptureOutputOf(func() { DisplayNames("Holly Clare Rees") })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
