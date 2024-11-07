// Extend the program in Exercise 2 by slicing the full name into 3 slices. Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines. [Slices] [Structures]

package assignment9

import (
	"fmt"
	"strings"
)

func Assignment9() {
	var fullName string
	fmt.Print("Enter full name (first middle last): ")
	fmt.Scanln(&fullName)
	DisplayNames(fullName)

}

func DisplayNames(fullName string) {
	names := strings.Split(fullName, " ")

	nameMap := map[string]string{
		"First name":  names[0],
		"Middle name": names[1],
		"Last name":   names[2],
	}

	keys := []string{"First name", "Middle name", "Last name"}

	for _, key := range keys {
		fmt.Printf("%s: %s\n", key, nameMap[key])
	}
}
