package recursion

import (
	"reflect"
	"testing"
)

func TestSumRecursion(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		expected int
	}{
		{"Single element", []int{42}, 42},
		{"All positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"All negative numbers", []int{-1, -2, -3, -4, -5}, -15},
		{"Mixed positive and negative numbers", []int{-1, 2, -3, 4, -5}, -3},
		{"Array with zeros", []int{0, 0, 0, 0}, 0},
		{"Large array", []int{100, 200, 300, 400, 500}, 1500},
		{"Empty array", []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SumRecursion(test.inputArr, len(test.inputArr)-1)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, result)
			}
		})
	}
}
