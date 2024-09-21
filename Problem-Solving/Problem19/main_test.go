package main

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := findMin(test.nums)
			if got != test.want {
				t.Errorf("expected %d but got %d", test.want, got)
			}
		})
	}
}
