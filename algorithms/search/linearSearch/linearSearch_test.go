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
