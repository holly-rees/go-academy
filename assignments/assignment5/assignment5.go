// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
// Three single digit numbers from one method.
// Three double digit numbers from a second method.
// Three triple digit numbers from a third method.
// Finally sum all methods into a final sum in the main program.
package assignment5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Assignment5() {

	userInput, err := validateInput(os.Args)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	adder := Adder{input: userInput}
	adder.GetSingleDigitNumbersFromInput()
	adder.GetDoubleDigitNumbersFromInput()
	adder.GetTripleDigitNumbersFromInput()
	adder.SumNumbers()

	fmt.Println("The total sum is:", adder.sum)

}

func validateInput(args []string) (string, error) {
	if len(args) != 2 || len(args[1]) != 9 {
		return "", fmt.Errorf("need a 9 digit number")
	}
	return args[1], nil
}

type Adder struct {
	input   string
	numbers []int
	sum     int
}

func (a *Adder) GetSingleDigitNumbersFromInput() {
	stringSlice := strings.Split(a.input[:3], "")
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("error converting to int")
			return
		}
		a.numbers = append(a.numbers, num)
	}
}

func (a *Adder) GetDoubleDigitNumbersFromInput() {
	for i := 0; i < 6; i += 2 {
		str := a.input[i : i+2]
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("error converting to int")
			return
		}
		a.numbers = append(a.numbers, num)
	}
}

func (a *Adder) GetTripleDigitNumbersFromInput() {
	for i := 0; i < 9; i += 3 {
		str := a.input[i : i+3]
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("error converting to int")
			return
		}
		a.numbers = append(a.numbers, num)
	}
}

func (a *Adder) SumNumbers() {
	for _, num := range a.numbers {
		a.sum += num
	}
}
