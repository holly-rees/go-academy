package assignment8

import (
	h "go-assignments/assignments"
	"os"
	"slices"
	"testing"
)

func TestCopyListToFile(t *testing.T) {
	fileName := "cities.txt"

	expected := "hello\nworld\n"

	CopyListToCitiesFile([]string{"hello", "world"})

	bytes, _ := os.ReadFile(fileName)

	actual := string(bytes)

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

}

func TestReadCitiesFromFile(t *testing.T) {
	expected := []string{"newcastle", "sunderland"}

	CopyListToCitiesFile([]string{"newcastle", "sunderland"})
	actual := ReadCitiesFromFile()

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestAssignment8(t *testing.T) {
	expected := "Abu Dhabi\nCaracas\nHanoi\nLondon\nMontevideo\nVatican City\nWashington D.C.\n"

	actual := h.CaptureOutputOf(Assignment8)

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
