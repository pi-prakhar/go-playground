package main

/**
PROBLEM:
leetcode -> Search in Rotated Sorted Array -> medium

TAGS:
array , binary search

**/

func search(nums []int, target int) int {
	k := searchMin(nums)
	return searchValue(nums, target, k)
}

func searchMin(nums []int) int {
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
	return index
}

func searchValue(nums []int, target int, k int) int {
	length := len(nums)
	l := 0
	h := length - 1

	for l <= h {
		mid := (l + h) / 2
		resolvedMid := (mid + k) % length
		if nums[resolvedMid] == target {
			return resolvedMid
		}
		if nums[resolvedMid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
