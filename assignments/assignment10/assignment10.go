// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]

package assignment10

import (
	"fmt"
	"time"
)

func Assignment10() {

}

type Pupil struct {
	Fullname    string
	DateOfBirth time.Time
	Age         int
}

type SchoolRegister struct {
	Pupils []Pupil
}

func (s SchoolRegister) ListPupils() {
	for _, pupil := range s.Pupils {
		fmt.Println(pupil.FormatAsString())
	}
}

func (p Pupil) FormatAsString() string {
	return fmt.Sprintf("Name: %s, Date of Birth: %s, Age: %d", p.Fullname, p.DateOfBirth.Format("2006-01-02"), p.Age)
}
