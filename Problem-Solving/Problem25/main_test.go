package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			input: []int{0},
			expected: [][]int{
				{0},
			},
		},
		{
			input: []int{1, 2},
			expected: [][]int{
				{1, 2}, {2, 1},
			},
		},
	}

	for _, test := range tests {
		result := permute(test.input)
		sortSubsets(result)
		sortSubsets(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func sortSubsets(subsets [][]int) {
	for _, subset := range subsets {
		sort.Ints(subset)
	}
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
}
