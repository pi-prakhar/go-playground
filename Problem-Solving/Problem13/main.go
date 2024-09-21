package main

/**
Problem
Leetcode ->  3Sum -> medium


Tags
arrays, 2 pointer
**/
import "sort"

func twoSum(res *[][]int, nums []int, target int, start int, end int) {
	iStart := start
	iEnd := end
	for iStart < iEnd {
		sum := nums[iStart] + nums[iEnd]
		if sum > target {
			iEnd--
			for iStart < iEnd && nums[iEnd] == nums[iEnd+1] {
				iEnd--
			}
		} else if sum == target {
			*res = append(*res, []int{0 - target, nums[iStart], nums[iEnd]})
			iStart++
			iEnd--
			for iStart < iEnd && nums[iStart] == nums[iStart-1] {
				iStart++
			}
			for iStart < iEnd && nums[iEnd] == nums[iEnd+1] {
				iEnd--
			}
		} else {
			iStart++
			for iStart < iEnd && nums[iStart] == nums[iStart-1] {
				iStart++
			}
		}

	}
}

func threeSum(nums []int) [][]int {
	var res [][]int
	end := len(nums) - 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for indexRoot := 0; indexRoot < len(nums)-2; indexRoot++ {
		numRoot := nums[indexRoot]
		if indexRoot > 0 && numRoot == nums[indexRoot-1] {
			continue
		}
		start := indexRoot + 1
		target := 0 - numRoot
		twoSum(&res, nums, target, start, end)
	}
	return res
}
