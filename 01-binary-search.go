package binarySearch

import (
	"errors"
	"fmt"
)

func binarySearch(values []int, search int, x int, y int) (error, int) {
	var mid int
	if len(values) == 0 {
		return errors.New("Empty array"), 0
	}

	if y == 0 {
		y = len(values) - 1
	}

	mid = (x + y) / 2

	if search == values[mid] {
		fmt.Printf("s=%v, Found=%v \n", search, mid)
		return nil, mid
	}

	if search < values[mid] {
		fmt.Printf("Lower, s=%v, x=%v, y=%v, mid=%v \n", search, x, mid, mid)
		return binarySearch(values, search, x, mid)
	}

	if search > values[mid] {
		fmt.Printf("Higher, s=%v,x=%v, y=%v, mid=%v \n", search, mid+1, y, mid)
		return binarySearch(values, search, mid+1, y)
	}

	if x == y {
		fmt.Printf("x==y, s=%v,x=%v, y=%v, mid=%v \n", search, x, y, mid)
		return nil, x
	}

	return nil, 0
}

func main() {

}
