// Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line. [Keyboard input]
package assignments

import "fmt"

func Assignment2() {
	var firstName, middleName, lastName string
	fmt.Print("Enter first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter middle name: ")
	fmt.Scanln(&middleName)
	fmt.Print("Enter last name: ")
	fmt.Scanln(&lastName)
	formattedStr := ConcatenateNames(firstName, middleName, lastName)
	fmt.Println("Your name is: " + formattedStr)
}

func ConcatenateNames(firstName, middleName, surname string) string {
	formattedString := fmt.Sprintf("%s %s %s", firstName, middleName, surname)
	return formattedString
}
