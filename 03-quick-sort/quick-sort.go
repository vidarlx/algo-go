package QuickSort

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var lower []int
	var greater []int

	var pivot int = arr[0]
	var i int = 1
	for i < len(arr) {
		if arr[i] < pivot {
			lower = append(lower, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
		i++
	}

	return append(append(quickSort(lower), pivot), quickSort(greater)...)
}

func main() {

}
