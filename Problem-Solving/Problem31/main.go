package main

/**
PROBLEM :
leetcode -> House Robber -> medium

TAGS:
dynamic programming, dp, array, notbest

**/
func solve(max *int, store *map[int]int, nums []int, start int) int {

	if val, ok := (*store)[start]; ok {
		return val
	}
	sum := 0
	tempmax := 0
	for index := start; index < len(nums); index++ {

		if index+2 < len(nums) {
			sum = nums[start] + solve(max, store, nums, index+2)
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
	solve(&max, &dp, nums, 0)
	solve(&max, &dp, nums, 1)
	return max
}
