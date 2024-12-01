package recursion

import "errors"

func MaxRecursion(arr []int, n int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("the array is empty")
	}

	if n == 0 {
		return arr[0], nil
	}

	max, err := MaxRecursion(arr, n-1)
	if err != nil {
		return 0, errors.New("something went wrong")
	}

	if max < arr[n] {
		return arr[n], nil
	} else {
		return max, nil
	}
}
