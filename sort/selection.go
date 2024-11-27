package sort

func SelectionSort(arr []int) []int {
	i := 0

	for ; i < len(arr)-1; i++ {
		min := i
		j := i + 1

		for ; j < len(arr); j++ {
			// Find min
			if arr[min] > arr[j] {
				min = j
			}
		}

		// Swap
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
