package bubbleSort

import "testing"

type testCase struct {
	name          string
	array         []int
	expectedArray []int
}

func TestBubbleSort(t *testing.T) {
	testCases := []testCase{
		{"Empty Array", []int{}, []int{}},
		{"Sorted Array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse Sorted Array", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Array with Duplicates", []int{3, 2, 3, 1, 2, 4, 4, 5, 5, 1}, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			BubbleSort(tc.array)

			if !arrayIsSorted(tc.array) && !arrayAreEquals(tc.array, tc.expectedArray) {
				t.Errorf("BubbleSort(arr), got : %v, want : %v", tc.array, tc.expectedArray)
			}
		})
	}
}

func arrayIsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func arrayAreEquals(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
