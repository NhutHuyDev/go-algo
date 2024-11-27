package sort

import (
	"reflect"
	"testing"
)

func HandleSortAgorithmTest(t *testing.T, sortFunc func(arr []int) []int) {
	tests := []struct {
		name     string
		inputArr []int
		expected []int
	}{
		{"Already sorted", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}},
		{"Reverse order", []int{88, 75, 67, 66, 56, 42, 32, 29, 13, 12, 11, 7, 3}, []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}},
		{"Unsorted array", []int{75, 67, 42, 11, 13, 88, 7, 32, 56, 29, 12, 3, 66}, []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}},
		{"Array with duplicates", []int{42, 42, 29, 13, 13, 88, 7, 56, 29, 12, 3, 66, 66}, []int{3, 7, 12, 13, 13, 29, 29, 42, 42, 56, 66, 66, 88}},
		{"Empty array", []int{}, []int{}},
		{"Single element", []int{42}, []int{42}},
		{"All elements identical", []int{7, 7, 7, 7}, []int{7, 7, 7, 7}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := sortFunc(test.inputArr)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, result)
			}
		})
	}
}
