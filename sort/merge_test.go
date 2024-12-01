package sort

import (
	"reflect"
	"testing"
)

func TestMergeFunc(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		left     int
		mid      int
		right    int
		expected []int
	}{
		{
			name:     "Basic Merge",
			arr:      []int{1, 3, 5, 2, 4, 6},
			left:     0,
			mid:      2,
			right:    5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Single Element",
			arr:      []int{1, 2},
			left:     0,
			mid:      0,
			right:    1,
			expected: []int{1, 2},
		},
		{
			name:     "Already Sorted",
			arr:      []int{1, 2, 3, 4},
			left:     0,
			mid:      1,
			right:    3,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.arr, tt.left, tt.mid, tt.right)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("got %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
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
			MergeSort(test.arr, 0, len(test.arr)-1)
			if !reflect.DeepEqual(test.arr, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.arr, test.expected, test.arr)
			}
		})
	}
}
