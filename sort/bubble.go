package sort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		isSwapped := false

		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	return arr
}
