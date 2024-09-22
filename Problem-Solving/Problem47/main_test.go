package main

import "testing"

func Test_largestNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want string
	}{
		{nums: []int{1, 2, 3, 4}, want: "4321"},
		{nums: []int{0}, want: "0"},
		{nums: []int{3, 30, 34}, want: "34330"},
		{nums: []int{9, 97}, want: "997"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := largestNumber(test.nums)
			if got != test.want {
				t.Errorf("wanted %s but got %s", test.want, got)
			}
		})
	}
}
