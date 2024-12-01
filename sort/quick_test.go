package sort

import (
	"reflect"
	"testing"
)

func TestQuickSortV1(t *testing.T) {
	HandleSortAgorithmTest(t, QuickSortV1)
}

func TestQuickSortV2PartitionFunc(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		expected int
	}{
		{"Already sorted", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, 12},
		{"Reverse order", []int{88, 75, 67, 66, 56, 42, 32, 29, 13, 12, 11, 7, 3}, 0},
		{"Unsorted array", []int{75, 67, 42, 11, 13, 88, 7, 32, 56, 29, 12, 3, 66}, 9},
		{"Array with duplicates", []int{42, 42, 29, 13, 13, 88, 7, 56, 29, 12, 3, 66, 66}, 11},
		{"Single element", []int{42}, 0},
		{"All elements identical", []int{7, 7, 7, 7}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Partition(test.inputArr, 0, len(test.inputArr)-1)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, result)
			}
		})
	}
}

func TestQuickSortV2(t *testing.T) {
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
			QuickSortV2(test.inputArr, 0, len(test.inputArr)-1)
			if !reflect.DeepEqual(test.inputArr, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, test.inputArr)
			}
		})
	}
}

func TestQuickSortV3(t *testing.T) {
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
			QuickSortV3(test.inputArr, 0, len(test.inputArr)-1)
			if !reflect.DeepEqual(test.inputArr, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, test.inputArr)
			}
		})
	}
}

func BenchmarkQuickSortV1(b *testing.B) {
	benchmarks := []struct {
		name     string
		inputArr []int
	}{
		{"Large random array", generateRandomArray(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := make([]int, len(bm.inputArr))
				copy(arr, bm.inputArr)
				QuickSortV1(arr)
			}
		})
	}
}

func BenchmarkQuickSortV2(b *testing.B) {
	benchmarks := []struct {
		name     string
		inputArr []int
	}{
		{"Large random array", generateRandomArray(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := make([]int, len(bm.inputArr))
				copy(arr, bm.inputArr)
				QuickSortV2(arr, 0, len(arr)-1)
			}
		})
	}
}
