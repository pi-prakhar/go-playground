package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func makeSliceMap(matrix [][]string) map[string]int {
	sliceMap := make(map[string]int)

	for _, row := range matrix {
		// Sort elements of the row
		sort.Strings(row)

		// Create a string representation of the sorted row to use as key
		key := fmt.Sprintf("%v", row)

		// Count occurrences of each sorted row
		sliceMap[key]++
	}

	return sliceMap
}

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		testName string
		strs     []string
		want     [][]string
	}{
		{
			"multiple strings",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"empty string",
			[]string{""},
			[][]string{{""}},
		},
		{
			"single string",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := groupAnagrams(test.strs)
			// Create maps of sorted slices
			map1 := makeSliceMap(got)
			map2 := makeSliceMap(test.want)
			if !reflect.DeepEqual(map1, map2) {
				t.Errorf("Expected %v but got %v", test.want, got)
			}
		})
	}
}
