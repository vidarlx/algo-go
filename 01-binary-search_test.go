package binarySearch

import "testing"

type testpair struct {
	values []int
	search int
	expect int
}

var tests = []testpair{
	{[]int{}, 1, 0},
	{[]int{1}, 1, 0},
	{[]int{1, 5}, 1, 0},
	{[]int{1, 5}, 5, 1},
	{[]int{0, 1, 12, 15, 18, 21, 25}, 0, 0},
	{[]int{0, 1, 12, 15, 18, 21, 25}, 15, 3},
	{[]int{0, 1, 12, 15, 18, 21, 25}, 25, 6},
}

func TestBinarySearch(t *testing.T) {
	var result int
	for _, test := range tests {
		_, result = binarySearch(test.values, test.search, 0, 0)
		if result != test.expect {
			t.Error(
				"For", test.values,
				"expected", test.expect,
				"got", result,
			)
		}
	}
}
