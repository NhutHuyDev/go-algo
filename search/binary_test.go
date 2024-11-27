package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		target   int
		expected int
	}{
		{"Target in middle", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, 32, 6},
		{"Target at start", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, 3, 0},
		{"Target at end", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, 88, 12},
		{"Target not found", []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}, 61, -1},
		{"Empty array", []int{}, 3, -1},
		{"Single element (found)", []int{3}, 3, 0},
		{"Single element (not found)", []int{3}, 4, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BinarySearch(test.inputArr, test.target)
			if result != test.expected {
				t.Errorf("For input %v and target %d, expected %d but got %d",
					test.inputArr, test.target, test.expected, result)
			}
		})
	}
}
