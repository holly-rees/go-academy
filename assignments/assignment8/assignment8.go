// Create a program that: [Write File][Read File][I/O Package][I/O]
// Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// Reads a list of cities from the newly created file.
// Displays the list of cities in alphabetical order.

package assignment8

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Assignment8() {
	cities := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

	CopyListToCitiesFile(cities)

	readCities := ReadCitiesFromFile()

	SortCitiesAlphabetically(readCities)

	DisplayCities(readCities)
}

func DisplayCities(readCities []string) {
	for _, city := range readCities {
		fmt.Println(city)
	}
}

func CopyListToCitiesFile(list []string) {

	file, _ := os.Create("cities.txt")
	defer file.Close()

	for _, city := range list {
		file.WriteString(city + "\n")
	}
}

func ReadCitiesFromFile() []string {
	bytes, _ := os.ReadFile("cities.txt")
	citiesStr := strings.TrimSpace(string(bytes))
	cities := strings.Split(citiesStr, "\n")
	return cities
}

func SortCitiesAlphabetically(cities []string) []string {
	sort.Strings(cities)
	return cities
}
