package main

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 1},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 11},
		{[]int{2, 1, 1, 2}, 3},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 1}, 2},
		{[]int{200, 3, 140, 20, 10}, 340},
		{[]int{6, 6, 4, 8, 4, 3, 3, 10}, 27},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := rob(test.nums)
			if got != test.want {
				t.Errorf("wanted %d but got %d", test.want, got)
			}
		})
	}
}
