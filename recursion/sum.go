package recursion

func SumRecursion(arr []int, n int) int {
	if len(arr) == 0 {
		return 0
	}

	if n == 0 {
		return arr[0]
	}

	return SumRecursion(arr, n-1) + arr[n]
}
