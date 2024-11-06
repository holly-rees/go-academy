package assignment5

import (
	"slices"
	"testing"
)

func TestGetSingleDigitNumbers(t *testing.T) {
	expected := []int{1, 2, 3}
	adder := Adder{input: "123456789"}

	adder.GetSingleDigitNumbersFromInput()
	actual := adder.numbers

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestGetDoubleDigitNumbers(t *testing.T) {
	expected := []int{12, 34, 56}
	adder := Adder{input: "123456789"}

	adder.GetDoubleDigitNumbersFromInput()
	actual := adder.numbers

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestGetTripleDigitNumbers(t *testing.T) {
	expected := []int{123, 456, 789}
	adder := Adder{input: "123456789"}

	adder.GetTripleDigitNumbersFromInput()
	actual := adder.numbers

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSumNumbers(t *testing.T) {
	expected := 6
	adder := Adder{input: "123456789"}

	adder.GetSingleDigitNumbersFromInput()
	adder.SumNumbers()

	actual := adder.sum
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
