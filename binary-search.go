package main

import (
	"fmt"
)

func binary_search(arr []int, n int) *int {
	start_idx := 0
	end_idx := len(arr)
	mid_idx := start_idx + (end_idx-start_idx)/2

	for mid_idx > 0 {
		if arr[mid_idx] == n {
			return &mid_idx
		} else if arr[mid_idx] > n {
			end_idx = mid_idx - 1
		} else {
			start_idx = mid_idx + 1
		}
		mid_idx = start_idx + (end_idx-start_idx)/2
	}

	return nil
}

func main() {
	arr := []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}

	fmt.Println(*binary_search(arr, 32))
}
