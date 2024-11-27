package search

func BinarySearch(arr []int, n int) int {
	start_idx := 0
	end_idx := len(arr)
	mid_idx := start_idx + (end_idx-start_idx)/2

	for start_idx < end_idx {
		if arr[mid_idx] == n {
			return mid_idx
		} else if arr[mid_idx] > n {
			end_idx = mid_idx - 1
		} else {
			start_idx = mid_idx + 1
		}
		mid_idx = start_idx + (end_idx-start_idx)/2
	}

	return -1
}
