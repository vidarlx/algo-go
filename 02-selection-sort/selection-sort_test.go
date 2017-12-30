package selectionSort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type testpair struct {
	list   []int
	expect []int
}

var tests = []testpair{
	{[]int{1}, []int{1}},
	{[]int{3, 4}, []int{3, 4}},
	{[]int{4, 3}, []int{3, 4}},
	{[]int{4, 3, 17, 1}, []int{1, 3, 4, 17}},
}

func TestSelectionSort(t *testing.T) {
	var result []int
	for _, test := range tests {
		result = selectionSort(test.list)
		if reflect.DeepEqual(result, test.expect) == false {
			t.Error(
				"For", test.list,
				"expected", test.expect,
				"got", result,
			)
		}
	}
}

func TestSelectionSortOnRandom(t *testing.T) {
	var randomInts = randomInts(1, 1000, 10000)
	// slices are passed by reference, so make a copy
	sorted := make([]int, len(randomInts))
	b := make([]int, len(randomInts))
	copy(b, randomInts)
	copy(sorted, randomInts)

	sort.Sort(sort.IntSlice(sorted))

	var result = selectionSort(b)
	if reflect.DeepEqual(result, sorted) == false {
		t.Error(
			"For", randomInts,
			"expected", sorted,
			"got", result,
		)
	}
}

func randomInts(min int, max int, amount int) []int {
	var res []int
	var bytes int
	i := 0
	for i < amount {
		bytes = min + rand.Intn(max)
		res = append(res, bytes)
		i++
	}
	return res
}
