package main

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "4x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			name: "1x1 matrix",
			input: [][]int{
				{1},
			},
			expected: []int{1},
		},
		{
			name: "3x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := spiralOrder(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("spiralOrder(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
