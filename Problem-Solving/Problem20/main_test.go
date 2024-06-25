package main

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 9, 7},
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
