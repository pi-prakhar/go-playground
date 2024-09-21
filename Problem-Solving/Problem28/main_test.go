package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		s    string
		want [][]string
	}{
		{
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := partition(test.s)
			sortSubsets(got)
			sortSubsets(test.want)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("For input %v, expected %v, but got %v", test.s, test.want, got)
			}
		})
	}
}

func sortSubsets(subsets [][]string) {
	// Sort individual subsets
	for _, subset := range subsets {
		sort.Strings(subset)
	}

	// Sort the list of subsets
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
}
