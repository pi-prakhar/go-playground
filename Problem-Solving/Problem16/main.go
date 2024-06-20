package main

/**
Problem:
leetcode -> Container With Most Water -> medium

Tags:
array, two pointer, math
**/
func maxArea(height []int) int {
	maxArea := 0
	start := 0
	end := len(height) - 1
	for start < end {
		heightStart := height[start]
		heightEnd := height[end]
		var currArea int

		if heightStart < heightEnd {
			currArea = heightStart * (end - start)
			start++

		} else {
			currArea = heightEnd * (end - start)
			end--
		}

		if currArea > maxArea {
			maxArea = currArea
		}
	}
	return maxArea
}
