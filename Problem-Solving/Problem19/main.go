package main

/**
PROBLEM:
leetcode -> Find Minimum in Rotated Sorted Array -> medium

Tags:
array, binary search
**/

func findMin(nums []int) int {
	l := 0
	h := len(nums) - 1
	mid := 0
	target := nums[l]
	index := l

	for l <= h {
		mid = (l + h) / 2
		if nums[mid] >= target {
			l = mid + 1
		} else {
			target = nums[mid]
			index = mid
			h = mid - 1
		}
	}

	return nums[index]
}

// func findMin(nums []int) int {
// 	l := len(nums)
// 	left := 0
// 	right := l - 1
// 	mid := (left + right) / 2
// 	gap := mid - left

// 	for !((nums[left] <= nums[mid]) && (nums[mid] <= nums[right])) {
// 		if (nums[mid] >= nums[left]) && (nums[left] >= nums[right]) {
// 			left = right
// 		} else {
// 			left = mid
// 		}
// 		right = (left + l - 1) % l
// 		mid = (left + gap) % l
// 	}
// 	return nums[left]
// }
