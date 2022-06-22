// From slice remove Nth element from the end.
// Из среза удалить N-й элемент с конца.
package main

import (
	"fmt"
)

// Result:
// Input:  [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(s, 3))
}

// Moving "down" one element position from higher numbers.
// Перенос "вниз" на одну позицию элементов с более высокими номерами.
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
