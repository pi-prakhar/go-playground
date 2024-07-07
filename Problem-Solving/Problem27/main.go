package main

import (
	"fmt"
	"sort"
)

/**
PROBLEM:
leetcode -> conmbinationSum2 -> medium

Tags:
recursion, arrays, backtracking
**/

func combinationSum2(candidates []int, target int) [][]int {
	var list [][]int
	sort.Slice(candidates, func(i int, j int) bool {
		return candidates[i] < candidates[j]
	})
	solve(&list, &candidates, []int{}, target, 0)
	return list

}

func solve(list *[][]int, nums *[]int, templist []int, target int, start int) {
	fmt.Println(templist)
	if target == 0 {
		copyOfTemp := make([]int, len(templist))
		copy(copyOfTemp, templist)
		(*list) = append((*list), copyOfTemp)
		return
	} else if target < 0 {
		return
	} else {
		for i := start; i < len((*nums)); i++ {
			if (i > start) && ((*nums)[i] == (*nums)[i-1]) {
				continue
			}
			templist = append(templist, (*nums)[i])
			solve(list, nums, templist, target-((*nums)[i]), i+1)
			templist = templist[:len(templist)-1]
		}
	}

}
