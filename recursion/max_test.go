package recursion

import (
	"errors"
	"reflect"
	"testing"
)

func TestMaxRecursion(t *testing.T) {
	tests := []struct {
		name      string
		inputArr  []int
		expected  int
		expectErr error
	}{
		{"Single element", []int{42}, 42, nil},
		{"Multiple elements", []int{-1, 22, 67, 13, 40}, 67, nil},
		{"Empty array", []int{}, 0, errors.New("the array is empty")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := MaxRecursion(test.inputArr, len(test.inputArr)-1)

			if test.expectErr != nil {
				if err.Error() != test.expectErr.Error() {
					t.Errorf("expected an error but got \"%v\" for input %v", err, test.inputArr)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got one for input %v: %v", test.inputArr, err)
				}
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("for input %v, expected %v but got %v", test.inputArr, test.expected, result)
				}
			}
		})
	}
}
