package main

import "sort"

/**
PROBLEM:
leetcode -> Subsets II -> medium

Tags:
recursion, arrays, backtracking
**/

func subsetsWithDup(nums []int) [][]int {
	var list [][]int
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})
	solve(&list, &nums, []int{}, 0)
	return list

}

func solve(list *[][]int, nums *[]int, templist []int, start int) {
	copyOfTemp := make([]int, len(templist))
	copy(copyOfTemp, templist)
	(*list) = append((*list), copyOfTemp)
	for i := start; i < len((*nums)); i++ {
		if (i > start) && (*nums)[i] == (*nums)[i-1] {
			continue
		}
		templist = append(templist, (*nums)[i])
		solve(list, nums, templist, i+1)
		templist = templist[:len(templist)-1]
	}

}
