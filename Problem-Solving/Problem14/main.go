package main

/**
Problem :
leetcode -> Product of Array Except Self -> medium

Tags :
arrays, math, algorithm
**/
func productExceptSelf(nums []int) []int {
	len := len(nums)
	right := make([]int, len)
	left := make([]int, len)
	prod := make([]int, len)

	left[0] = 1
	right[len-1] = 1

	for index := 1; index < len; index++ {
		left[index] = left[index-1] * nums[index-1]
	}
	for index := len - 2; index >= 0; index-- {
		right[index] = right[index+1] * nums[index+1]
	}
	for index := 0; index < len; index++ {
		prod[index] = left[index] * right[index]
	}

	return prod
}
