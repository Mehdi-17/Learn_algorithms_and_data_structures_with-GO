package binarySearch

import (
	"testing"
)

type testCase struct {
	name         string
	orderedArray []int
	searching    int
	expected     int
}

func TestBinarySearch(t *testing.T) {
	testCases := []testCase{
		{"ElementExists", []int{1, 2, 4, 5, 7, 12, 14, 15, 17, 27, 37, 117, 1717}, 17, 8},
		{"ElementExists", []int{1, 5, 7, 12, 18, 27, 37, 117, 1717}, 17, -1},
		{"MiddleInSortedArray", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{"FirstElement", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{"LastElement", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{"EvenNumbers", []int{2, 4, 6, 8, 10, 12, 14, 16}, 8, 3},
		{"MultiplesOf3", []int{3, 6, 9, 12, 15, 18, 21}, 18, 5},
		{"ElementBiggerThanWholeArray", []int{1, 2, 3, 4, 5}, 6, -1},
		{"ElementSmallerThanWholeArray", []int{2, 4, 6, 8, 10}, 1, -1},
		{"EmptyArray", []int{}, 5, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(tc.orderedArray, tc.searching)

			if result != tc.expected {
				t.Errorf("BinarySearch(arr, %v), got : %v, want : %v", tc.searching, result, tc.expected)
			}
		})
	}
}
