package main

func maxArea(height []int) int {
	maxArea := 0
	start := 0
	end := len(height) - 1
	for start < end {
		heightStart := height[start]
		heightEnd := height[end]
		var minHeight int

		if heightStart < heightEnd {
			start++
			minHeight = heightStart
		} else {
			end--
			minHeight = heightEnd
		}
		currArea := minHeight * (end - start)
		if currArea > maxArea {
			maxArea = currArea
		}
	}
	return maxArea

}
