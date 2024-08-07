package main

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Simple case", []int{1, 2, 3}, []int{1, 3, 2}},
		{"Reverse order", []int{3, 2, 1}, []int{1, 2, 3}},
		{"Duplicate numbers", []int{1, 1, 5}, []int{1, 5, 1}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements, ascending", []int{1, 2}, []int{2, 1}},
		{"Two elements, descending", []int{2, 1}, []int{1, 2}},
		{"Larger sequence", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 5, 4}},
		{"Larger sequence with swap", []int{1, 3, 5, 4, 2}, []int{1, 4, 2, 3, 5}},
		{"All same numbers", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"Empty slice", []int{}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			nextPermutation(input)
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("nextPermutation(%v) = %v; want %v", tc.input, input, tc.expected)
			}
		})
	}
}

func TestReverseNums(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		start    int
		end      int
		expected []int
	}{
		{"Reverse entire slice", []int{1, 2, 3, 4, 5}, -1, 5, []int{5, 4, 3, 2, 1}},
		{"Reverse partial slice", []int{1, 2, 3, 4, 5}, 1, 4, []int{1, 4, 3, 2, 5}},
		{"Reverse single element", []int{1}, -1, 1, []int{1}},
		{"Reverse empty slice", []int{}, -1, 0, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reverseNums(&tc.input, &tc.start, &tc.end)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("reverseNums(%v, %d, %d) = %v; want %v", tc.input, tc.start, tc.end, tc.input, tc.expected)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		start    int
		end      int
		expected []int
	}{
		{"Swap adjacent elements", []int{1, 2, 3}, 0, 1, []int{2, 1, 3}},
		{"Swap distant elements", []int{1, 2, 3, 4, 5}, 1, 4, []int{1, 5, 3, 4, 2}},
		{"Swap same element", []int{1, 2, 3}, 1, 1, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			swap(&tc.input, &tc.start, &tc.end)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("swap(%v, %d, %d) = %v; want %v", tc.input, tc.start, tc.end, tc.input, tc.expected)
			}
		})
	}
}

func TestLarger(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		start    int
		expected int
	}{
		{"Simple case", []int{1, 3, 5, 4, 2}, 3, 1, 3},
		{"All larger", []int{5, 6, 7, 8}, 4, 0, 0},
		{"Only last larger", []int{1, 2, 3, 4, 5}, 4, 0, 4},
		{"No larger", []int{5, 4, 3, 2, 1}, 5, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := larger(&tc.nums, &tc.target, &tc.start)
			if result != tc.expected {
				t.Errorf("larger(%v, %d, %d) = %d; want %d", tc.nums, tc.target, tc.start, result, tc.expected)
			}
		})
	}
}
