package main

import (
	"fmt"

	"github.com/NhutHuyDev/go-algo/sort"
)

func main() {
	// ## Binary Search
	// arr := []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}
	// fmt.Println(search.Binary_Search(arr, 32))

	// ## Selection Sort
	arr := []int{75, 67, 42, 11, 13, 88, 7, 32, 56, 29, 12, 3, 66}
	fmt.Println(sort.Selection_Sort(arr))
}
