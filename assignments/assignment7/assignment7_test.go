package assignment7

import (
	h "go-assignments/assignments"
	"testing"
)

type MockDice struct {
	Rolls []int
}

func (d MockDice) Roll() (int, int) {
	return d.Rolls[0], d.Rolls[1]
}

func TestNaturalDiceRoll7(t *testing.T) {
	mockDice := MockDice{Rolls: []int{1, 6}}

	expected := "Rolled a 1 and a 6: NATURAL!\n"

	actual := h.CaptureOutputOf(func() { RollDice(mockDice) })

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestNaturalDiceRoll11(t *testing.T) {
	mockDice := MockDice{Rolls: []int{5, 6}}

	expected := "Rolled a 5 and a 6: NATURAL!\n"

	actual := h.CaptureOutputOf(func() { RollDice(mockDice) })

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSnakeEyesRoll2(t *testing.T) {
	mockDice := MockDice{Rolls: []int{1, 1}}

	expected := "Rolled a 1 and a 1: SNAKE-EYES-CRAPS!\n"

	actual := h.CaptureOutputOf(func() { RollDice(mockDice) })

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestLossCrapsRoll3(t *testing.T) {
	mockDice := MockDice{Rolls: []int{2, 1}}

	expected := "Rolled a 2 and a 1: LOSS-CRAPS!\n"

	actual := h.CaptureOutputOf(func() { RollDice(mockDice) })

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestLossCrapsRoll12(t *testing.T) {
	mockDice := MockDice{Rolls: []int{6, 6}}

	expected := "Rolled a 6 and a 6: LOSS-CRAPS!\n"

	actual := h.CaptureOutputOf(func() { RollDice(mockDice) })

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
