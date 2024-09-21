package main

import (
	"reflect"
	"testing"
)

// Helper function to compare two slices of slices of integers
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func Test_threeSum(t *testing.T) {
	tests := []struct {
		testName string
		nums     []int
		want     [][]int
	}{
		{"multiple values", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"returns 0", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"no sum can be created", []int{1, 2}, [][]int{}},
		{"empty array", []int{}, [][]int{}},
		{"all positive", []int{1, 2, 3, 4, 5}, [][]int{}},
		{"all negative", []int{-1, -2, -3, -4, -5}, [][]int{}},
		{"all zero", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"repeated triplets", []int{-1, -1, 2, 2}, [][]int{{-1, -1, 2}}},
		{"large numnbers", []int{1000000000, -1000000000, 0, 1, -1}, [][]int{{-1000000000, 0, 1000000000}, {-1, 0, 1}}},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := threeSum(test.nums)
			if !equal(got, test.want) {
				t.Errorf("expected %v but got %v", test.want, got)
			}
		})
	}
}
