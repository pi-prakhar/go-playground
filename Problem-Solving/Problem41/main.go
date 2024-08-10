package main

/**
PROBLEM:
leetcode -> maxsubarray -> medium

TAGS:
array
**/
func maxSubArray(nums []int) int {
	tempSum := nums[0]
	maxSum := nums[0]

	for index := 1; index < len(nums); index++ {
		tempSum = tempSum + nums[index]
		if nums[index] > tempSum {
			tempSum = nums[index]
		}
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}

	return maxSum
}
