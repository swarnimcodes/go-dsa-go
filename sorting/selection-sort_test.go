package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 2, 9, 3, 1},
			expected: []int{1, 2, 3, 5, 9},
		},
		{
			input:    []int{10, 5, 8, 3, 6},
			expected: []int{3, 5, 6, 8, 10},
		},
	}

	for _, tc := range testCases {
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		selectionSort(tc.input)

		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("SelectionSort(%v) = %v, expected %v", inputCopy, tc.input, tc.expected)
		}
	}
}
