package main

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	testCases := []struct {
		input  []int
		answer []int
	}{
		{[]int{9, 7, 8, 5, 6, 4, 1, 2, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{5, 2, 6, 4, 2, 5}, []int{2, 2, 4, 5, 5, 6}},
		{[]int{1}, []int{1}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 5, 4, 5}, []int{4, 4, 5, 5, 5}},
		{[]int{}, []int{}},
	}

	for _, data := range testCases {
		copyInp := make([]int, len(data.input))
		copy(copyInp, data.input)

		Sorting(copyInp)

		if !reflect.DeepEqual(copyInp, data.answer) {
			t.Errorf("Your answer: %d, expected: %d", copyInp, data.answer)
		}
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		arr1   []int
		arr2   []int
		answer bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, true},  // Arrays are equal after sorting
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},          // Different lengths
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 6}, false}, // Arrays have different elements
		{[]int{}, []int{}, true},                            // Empty arrays are equal
	}

	for _, data := range testCases {
		actual := compare(data.arr1, data.arr2)
		if actual != data.answer {
			t.Errorf("Your answer: %v, expected: %v", actual, data.answer)
		}
	}
}
