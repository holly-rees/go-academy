// Create a program that has multiple string variable and displays the string on one line. [Strings]

package assignments

import "fmt"

func PrintHelloWorld() {
	greeting, subject, punctuation := "hello", "world", "!"
	formattedString := fmt.Sprintf("%s %s%s", greeting, subject, punctuation)
	fmt.Println(formattedString)
	// fmt.Println(greeting + " " + subject + punctuation)
}
