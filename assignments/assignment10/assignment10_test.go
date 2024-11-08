package assignment10

import (
	h "go-assignments/assignments"
	"testing"
	"time"
)

func TestSchoolRegisterInitsWithPupils(t *testing.T) {
	pupils := []Pupil{
		{"John Doe", time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC), 14},
		{"Jane Smith", time.Date(2011, time.February, 2, 0, 0, 0, 0, time.UTC), 13},
	}
	schoolRegister := SchoolRegister{Pupils: pupils}

	if len(schoolRegister.Pupils) != 2 {
		t.Errorf("Expected 2 pupils but got %d", len(schoolRegister.Pupils))
	}
}

func TestListPupilsMethod(t *testing.T) {
	pupils := []Pupil{
		{"John Doe", time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC), 14},
		{"Jane Smith", time.Date(2011, time.February, 2, 0, 0, 0, 0, time.UTC), 13},
	}
	schoolRegister := SchoolRegister{Pupils: pupils}

	expected := "Name: John Doe, Date of Birth: 2010-01-01, Age: 14\nName: Jane Smith, Date of Birth: 2011-02-02, Age: 13\n"
	actual := h.CaptureOutputOf(schoolRegister.ListPupils)

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestPupilFormatAsStringMethod(t *testing.T) {
	john := Pupil{"John Doe", time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC), 14}

	expected := "Name: John Doe, Date of Birth: 2010-01-01, Age: 14"
	actual := john.FormatAsString()

	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
