// Create a program that calculates the age of a person given their date of birth. [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
// (Use the github.com/bearbin/go-age to aid in the creation of this app. Also review unit testing applied to the application age.go within the imported package.)

package assignment6

import (
	"time"

	b "github.com/bearbin/go-age"
)

func Assignment6() {
	CalculateAge("13/09/1992")
}

func CalculateAge(dob string) int {
	layout := "02/01/2006"
	birthDate, _ := time.Parse(layout, dob)

	return b.Age(birthDate)
}
