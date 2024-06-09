package main

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 1}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{3, 3}, true},
		{[]int{3, 4, 5, 6, 3, 7, 8}, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("containsDuplicate(%v)", tt.nums), func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %t; want = %t", tt.nums, got, tt.want)
			}
		})

	}
}
