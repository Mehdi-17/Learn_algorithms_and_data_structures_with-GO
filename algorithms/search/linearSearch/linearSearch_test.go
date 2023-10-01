package linearSearch

import "testing"

type testCase struct {
	name      string
	array     []int
	searching int
	expected  int
}

func TestLinearSearch(t *testing.T) {
	testCases := []testCase{
		{"ElementExists", []int{15, 6, 17, 56, 561, 45}, 17, 2},
		{"ElementNoExists", []int{15, 6, 17, 56, 561, 45}, 196, -1},
		{"ElementExistsMiddle", []int{15, 6, 17, 56, 561, 45}, 17, 2},
		{"ElementExistsBeginning", []int{15, 6, 17, 56, 561, 45}, 15, 0},
		{"ElementExistsEnd", []int{15, 6, 17, 56, 561, 45}, 45, 5},
		{"ElementExistsSmallArray", []int{1, 2, 3, 4, 5}, 5, 4},
		{"ElementExistsSmallArray", []int{1, 2, 3, 4, 5}, 2, 1},
		{"ElementNotExistsSmallArray", []int{1, 2, 3, 4, 5}, 6, -1},
		{"ElementNotExistsSingleElementArray", []int{1}, 2, -1},
		{"EmptyArray", []int{}, 5, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LinearSearch(tc.array, tc.searching)
			if result != tc.expected {
				t.Errorf("linearSearch(arr, %v), got : %v, want : %v,  ", tc.searching, result, tc.expected)
			}
		})
	}
}
