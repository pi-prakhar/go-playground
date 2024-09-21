package main

import "sort"

/**
PROBLEM:
leetcode -> Merge Intervals -> medium

Tags:
array, sort
**/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	min := intervals[0][0]
	max := intervals[0][1]
	for index := 1; index < len(intervals); index++ {
		start := intervals[index][0]
		end := intervals[index][1]

		if start > max {
			result = append(result, []int{min, max})
			min = start
			max = end
			continue
		}

		if (start <= max) && (end > max) {
			max = end
		}
	}

	result = append(result, []int{min, max})

	return result
}
