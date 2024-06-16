package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		testName string
		nums     []int
		k        int
		want     []int
	}{
		{"Many element", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"One element", []int{1}, 1, []int{1}},
		{"All order", []int{1, 1, 1, 3, 4, 2, 1, 3, 3, 2, 4, 2}, 4, []int{1, 3, 2, 4}},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := topKFrequent(test.nums, test.k)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expected %v but got %v", test.want, got)
			}
		})
	}
}
