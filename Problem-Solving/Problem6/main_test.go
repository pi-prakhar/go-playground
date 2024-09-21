package main

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{7, 1, 8, 0, 2, 3, 4}, 7},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("maxProfit{%v}", tt.prices), func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit({%v}) = %d ; want {%d}", tt.prices, got, tt.want)
			}
		})
	}
}
