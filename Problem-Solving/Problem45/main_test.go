package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Basic case",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "nums1 is empty",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums2 is empty",
			nums1:    []int{1, 2, 3},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both arrays have same elements",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "nums1 has larger elements",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "nums2 has larger elements",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Single element arrays",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "Large arrays",
			nums1:    []int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0},
			m:        5,
			nums2:    []int{2, 4, 6, 8, 10},
			n:        5,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Arrays with negative numbers",
			nums1:    []int{-3, -1, 0, 0, 0},
			m:        2,
			nums2:    []int{-2, -1, 1},
			n:        3,
			expected: []int{-3, -2, -1, -1, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			if !reflect.DeepEqual(tc.nums1, tc.expected) {
				t.Errorf("merge(%v, %d, %v, %d) = %v; want %v", tc.nums1, tc.m, tc.nums2, tc.n, tc.nums1, tc.expected)
			}
		})
	}
}
