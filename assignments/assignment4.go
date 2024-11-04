// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// Display the array content in ascending sequential order 1 to 10.
// Display the array content in descending sequential order 10 to 1.
// Count even numbers and odd numbers in increasing and decreasing sequential order.
// Display the even and odd count sequences to screen.

package assignments

import (
	"fmt"
	"sort"
)

func Assignment4() {

}

func Create1to10Array() []int {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return slice
}

func DisplayAscendingSlice(slice []int) {
	sort.Ints(slice)
	fmt.Println(slice)
}

func DisplayDescendingSlice(slice []int) {
	sort.Ints(slice)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	fmt.Println(slice)
}
