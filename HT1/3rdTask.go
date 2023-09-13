package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 3, 4, 2}
	arr3 := []int{1, 2, 3, 4, 5}

	fmt.Println(compare(arr1, arr2))
	fmt.Println(compare(arr1, arr3))
}
func compare(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	sort1 := make([]int, len(arr1))
	copy(sort1, arr1)
	Sorting(sort1)

	sort2 := make([]int, len(arr2))
	copy(sort2, arr2)
	Sorting(sort2)

	return reflect.DeepEqual(sort1, sort2)
}
