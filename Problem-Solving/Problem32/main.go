package main

// /**
// PROBLEM :
// leetcode -> House Robber 2-> medium

// TAGS:
// dynamic programming, dp, array , notbest

// **/

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
	dp2 := make(map[int]int)
	newNums := nums[0 : len(nums)-1]
	solve(&max, &dp, newNums, 0)
	solve(&max, &dp2, nums, 1)
	solve(&max, &dp2, nums, 2)

	return max
}
