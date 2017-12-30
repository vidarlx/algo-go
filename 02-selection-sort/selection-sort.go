package selectionSort

func findMin(arr []int) int {
	var min int
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return min
}

func selectionSort(arr []int) []int {
	var min int
	var result []int

	if len(arr) == 1 {
		return arr
	}

	for 0 < len(arr) {
		min = findMin(arr)
		result = append(result, arr[min])
		arr = append(arr[:min], arr[min+1:]...)
	}

	return result
}

func main() {

}
