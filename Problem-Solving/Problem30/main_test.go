package main

import "testing"

func Test_climbstairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"with input 2", 2, 2},
		{"with input 0", 0, 1},
		{"with input 1", 1, 1},
		{"with input 45", 45, 1836311903},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := climbStairs(test.n)
			if test.want != got {
				t.Errorf("wanted %d but got %d", test.want, got)
			}
		})
	}
}
