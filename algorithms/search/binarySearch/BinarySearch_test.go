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
