package main

import (
	"fmt"
	//"reflect"
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

	aMap := make(map[int]int)
	bMap := make(map[int]int)

	for i := 0; i < len(arr1); i++ {
		aMap[arr1[i]]++
	}

	for i := 0; i < len(arr2); i++ {
		bMap[arr2[i]]++
	}

	for aVal, aKey := range aMap {
		if bMap[aVal] != aKey {
			return false
		}
	}

	return true
}
