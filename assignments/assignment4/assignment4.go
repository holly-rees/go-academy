// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// Display the array content in ascending sequential order 1 to 10.
// Display the array content in descending sequential order 10 to 1.
// Count even numbers and odd numbers in increasing and decreasing sequential order.
// Display the even and odd count sequences to screen.

package assignment4

import (
	"fmt"
	"sort"
	"strings"
)

func Assignment4() {
	slice := Create1to10Slice()
	fmt.Printf("slice: %v \n", slice)
	fmt.Println("Ascending odd and even: ")
	DisplayOddNumbers(slice)
	DisplayEvenNumbers(slice)
	fmt.Println("Descending odd and even: ")
	DisplayOddNumbers(slice, "descending")
	DisplayEvenNumbers(slice, "descending")
}

func Create1to10Slice() []int {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return slice
}

func DisplayAscendingSlice(slice []int) {
	sort.Ints(slice)
	fmt.Println(slice)
}

func DisplayDescendingSlice(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	fmt.Println(slice)
}

// variadic input parameters ...string allows more than one (or none)
func DisplaySliceInDirection(slice []int, direction ...string) {
	if len(direction) > 0 && strings.ToLower(direction[0]) == "descending" {
		DisplayDescendingSlice(slice)
	} else {
		DisplayAscendingSlice(slice)
	}
}

func DisplayEvenNumbers(slice []int, direction ...string) {
	even := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	DisplaySliceInDirection(even, direction...)
}

func DisplayOddNumbers(slice []int, direction ...string) {
	odd := []int{}
	for _, num := range slice {
		if num%2 != 0 {
			odd = append(odd, num)
		}
	}
	DisplaySliceInDirection(odd, direction...)
}
