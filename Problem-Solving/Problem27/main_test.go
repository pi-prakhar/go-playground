package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected [][]int
	}{
		{
			input:  []int{2, 3, 6, 7},
			target: 7,
			expected: [][]int{
				{7},
			},
		},
		{
			input:  []int{2, 3, 5},
			target: 8,
			expected: [][]int{
				{3, 5},
			},
		},
		{
			input:  []int{10, 1, 2, 7, 6, 1, 5},
			target: 8,
			expected: [][]int{
				{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
			},
		},
		{
			input:  []int{2, 5, 2, 1, 2},
			target: 5,
			expected: [][]int{
				{1, 2, 2}, {5},
			},
		},
	}

	for _, test := range tests {
		result := combinationSum2(test.input, test.target)
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
