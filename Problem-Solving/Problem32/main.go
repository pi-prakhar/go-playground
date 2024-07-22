package main

/**
PROBLEM :
leetcode -> House Robber 2-> medium

TAGS:
dynamic programming, dp, array

**/
func solve(max *int, store *map[int]int, nums []int, start int, skip bool) int {

	if val, ok := (*store)[start]; ok {
		return val
	}
	sum := 0
	tempmax := 0
	for index := start; index < len(nums); index++ {
		if skip && index == len(nums)-1 {
			break
		} else if index+2 < len(nums) {
			sum = nums[start] + solve(max, store, nums, index+2, skip)
		} else {
			sum = nums[start]
		}

		if sum > tempmax {
			tempmax = sum
		}
	}
	(*store)[start] = tempmax
	if tempmax > *max {
		*max = tempmax
	}
	return tempmax

}

func rob(nums []int) int {
	max := 0
	dp := make(map[int]int)
	solve(&max, &dp, nums, 0, true)
	solve(&max, &dp, nums, 1, false)
	solve(&max, &dp, nums, 2, false)
	return max
}
