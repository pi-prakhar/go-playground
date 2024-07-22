package main

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		// {[]int{1, 2, 3, 1}, 4},
		// {[]int{2, 7, 9, 3, 1}, 11},
		// {[]int{2, 1, 1, 2}, 3},
		{[]int{1, 2, 3}, 3},
		// {[]int{1, 2, 1}, 2},
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
