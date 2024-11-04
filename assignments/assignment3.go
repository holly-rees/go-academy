// Create a program that allows the user to input a number. Check whether the number lies between 1 and 10. [Variables]
package assignments

import "fmt"

func Assignment3() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	if isBetween1and10(number) {
		fmt.Println("It's between 1 and 10!")
	} else {
		fmt.Println("It's NOT between 1 and 10!")
	}
}

func isBetween1and10(number int) bool {
	return number < 10 && number > 1
}
