package recursion

import (
	"reflect"
	"testing"
)

func TestCountRecursion(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		expected int
	}{
		{"Single element", []int{42}, 1},
		{"Multiple elements", []int{-1, 22, 67, 13, 40}, 5},
		{"Empty array", []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountRecursion(test.inputArr, len(test.inputArr)-1)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v, expected %v but got %v", test.inputArr, test.expected, result)
			}
		})
	}
}
