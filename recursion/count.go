package recursion

func CountRecursion(arr []int, n int) int {
	if len(arr) == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return CountRecursion(arr, n-1) + 1
}
