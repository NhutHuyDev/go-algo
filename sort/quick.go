package sort

/* QuickSortV1 Implementation */

func QuickSortV1(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	pivot := arr[0]

	left := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		}
	}

	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] >= pivot {
			right = append(right, arr[i])
		}
	}

	sorted_arr := make([]int, 0)

	sorted_arr = append(sorted_arr, QuickSortV1(left)...)

	sorted_arr = append(sorted_arr, pivot)

	sorted_arr = append(sorted_arr, QuickSortV1(right)...)

	return sorted_arr
}

/* QuickSortV2 Implementation */

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[high], arr[i+1] = arr[i+1], arr[high]

	return i + 1
}

func QuickSortV2(arr []int, low, high int) {
	if low < high {
		pivot := Partition(arr, low, high)

		QuickSortV2(arr, low, pivot-1)
		QuickSortV2(arr, pivot+1, high)
	}
}
