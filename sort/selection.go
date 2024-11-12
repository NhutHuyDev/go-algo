package sort

func Selection_Sort(arr []int) []int {
	i := 0

	for i < len(arr)-1 {
		min := i
		j := i + 1

		for j < len(arr) {
			// Find min
			if arr[min] > arr[j] {
				min = j
			}
			j++
		}

		// Swap
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp

		i++
	}

	return arr
}
