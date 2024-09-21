package main

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := maxArea(test.height)
			if got != test.want {
				t.Errorf("expected %d but got %d", test.want, got)
			}
		})
	}
}
