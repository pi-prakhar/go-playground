package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"multiple numbers", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"containing zeroes", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := productExceptSelf(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expected %v but got %v", test.want, got)
			}
		})
	}
}
