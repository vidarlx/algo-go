package QuickSort

import (
	"reflect"
	"testing"
)

type testpair struct {
	list   []int
	expect []int
}

var tests = []testpair{
	{[]int{}, []int{}},
	{[]int{4}, []int{4}},
	{[]int{1, 4}, []int{1, 4}},
	{[]int{4, 1}, []int{1, 4}},
	{[]int{4, 1, 8, 3, 12}, []int{1, 3, 4, 8, 12}},
}

func TestQuickSort(t *testing.T) {
	var result []int
	for _, test := range tests {
		result = quickSort(test.list)
		if reflect.DeepEqual(result, test.expect) == false {
			t.Error(
				"For", test.list,
				"expected", test.expect,
				"got", result,
			)
		}
	}
}
