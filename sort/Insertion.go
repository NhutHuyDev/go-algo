package sort

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		key := arr[i+1]

		j := i
		for ; j >= 0; j-- {
			if key < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		arr[j+1] = key
	}
	return arr
}
