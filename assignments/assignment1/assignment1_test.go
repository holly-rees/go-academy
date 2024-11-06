package assignments

import (
	h "go-assignments/assignments"
	"testing"
)

func TestAssignment1(t *testing.T) {
	expected := "hello world!\n"
	actual := h.CaptureOutputOf(PrintHelloWorld)
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
