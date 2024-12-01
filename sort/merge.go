package sort

func Merge(arr []int, left, mid, right int) {
	n_left := mid - left + 1
	n_right := right - mid

	left_arr := make([]int, n_left)
	right_arr := make([]int, n_right)

	for i := 0; i < len(left_arr); i++ {
		left_arr[i] = arr[i+left]
	}

	for i := 0; i < len(right_arr); i++ {
		right_arr[i] = arr[mid+1+i]
	}

	var i, j, k int
	k = left
	for i < n_left && j < n_right {
		if left_arr[i] < right_arr[j] {
			arr[k] = left_arr[i]
			i++
		} else {
			arr[k] = right_arr[j]
			j++
		}
		k++
	}

	for i < n_left {
		arr[k] = left_arr[i]
		i++
		k++
	}

	for i < n_left {
		arr[k] = right_arr[j]
		j++
		k++
	}
}

func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (right + left) / 2

		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)

		Merge(arr, left, mid, right)
	}
}
