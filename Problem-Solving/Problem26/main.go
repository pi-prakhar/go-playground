package main

import "sort"

/**
PROBLEM:
leetcode -> conmbination sum -> medium

Tags:
recursion, arrays, backtracking
**/

func combinationSum(candidates []int, target int) [][]int {
	var list [][]int
	sort.Slice(candidates, func(i int, j int) bool {
		return candidates[i] < candidates[j]
	})
	solve(&list, &candidates, []int{}, 0, target, 0)
	return list

}

func solve(list *[][]int, nums *[]int, templist []int, sum int, target int, start int) {
	if target == sum {
		copyOfTemp := make([]int, len(templist))
		copy(copyOfTemp, templist)
		(*list) = append((*list), copyOfTemp)
		return
	} else if sum > target {
		return
	} else {
		for i := start; i < len((*nums)); i++ {
			templist = append(templist, (*nums)[i])
			sum += (*nums)[i]
			solve(list, nums, templist, sum, target, start)
			templist = templist[:len(templist)-1]
			sum -= (*nums)[i]
			start++
		}

	}

}
