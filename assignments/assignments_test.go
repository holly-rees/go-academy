package assignments

import (
	"bytes"
	"os"
	"slices"
	"testing"
)

func TestAssignment1(t *testing.T) {
	expected := "hello world!\n"
	actual := captureOutputOf(PrintHelloWorld)
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestConcatenateNames(t *testing.T) {
	expected := "Holly Clare Rees"
	actual := ConcatenateNames("Holly", "Clare", "Rees")
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestIsBetween1and10(t *testing.T) {
	expected := true
	actual := isBetween1and10(5)
	if actual != expected {
		t.Errorf("Expected %t but got %t", expected, actual)
	}
}

func TestCreateArray(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := Create1to10Slice()
	if !slices.Equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestDisplayAscendingSlice(t *testing.T) {
	slice := []int{1, 5, 2, 7}
	expected := "[1 2 5 7]\n"
	actual := captureOutputOf(func() { DisplayAscendingSlice(slice) })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestDisplayDescendingSlice(t *testing.T) {
	slice := []int{1, 5, 2, 7}
	expected := "[7 5 2 1]\n"
	actual := captureOutputOf(func() { DisplayDescendingSlice(slice) })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCountEvenNumbersAscending(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := "[2 4 6 8 10]\n"
	actual := captureOutputOf(func() { DisplayEvenNumbers(slice) })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCountEvenNumbersDescending(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := "[10 8 6 4 2]\n"
	actual := captureOutputOf(func() { DisplayEvenNumbers(slice, "descending") })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCountOddNumbersAscending(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := "[1 3 5 7 9]\n"
	actual := captureOutputOf(func() { DisplayOddNumbers(slice) })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestCountOddNumbersDescending(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := "[9 7 5 3 1]\n"
	actual := captureOutputOf(func() { DisplayOddNumbers(slice, "descending") })
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func captureOutputOf(function func()) string {
	var buffer bytes.Buffer
	originalOutputSetting := os.Stdout
	readEnd, writeEnd, _ := os.Pipe()
	os.Stdout = writeEnd //temporarily set the os standard out to be the write end of our pipe

	function()

	writeEnd.Close()
	os.Stdout = originalOutputSetting // reset it

	buffer.ReadFrom(readEnd)
	return buffer.String()
}
