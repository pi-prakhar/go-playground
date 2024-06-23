package main

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 4, -1},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := search(test.nums, test.target)
			if got != test.want {
				t.Errorf("expected %d but got %d", test.want, got)
			}
		})
	}
}
