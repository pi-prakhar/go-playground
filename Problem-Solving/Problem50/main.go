package main

/*
PROBLEM:
42. Trapping Rain Water, leetcode, Hard

TAGS:
double pointer, array, DP
*/

func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}
	start := 1
	end := length - 1
	totalWater := 0
	maxLeft := height[0]
	maxRight := height[2]
	indexR := 0

	for start < end {
		if height[start-1] > maxLeft {
			maxLeft = height[start-1]
		}
		if start >= indexR {
			indexR, maxRight = getMaxRight(start, height)
		}
		water := min(maxLeft, maxRight) - height[start]

		if water > 0 {
			totalWater += water
		}
		start += 1
	}
	return totalWater
}

func getMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMaxRight(index int, array []int) (int, int) {
	length := len(array)
	max := 0
	maxI := index
	for index < length-1 {
		index = index + 1
		if array[index] > max {
			max = array[index]
			maxI = index
		}
	}
	return maxI, max
}

func getMaxLeft(index int, array []int) (int, int) {
	max := 0
	maxI := index
	for index > 0 {
		index = index - 1
		if array[index] > max {
			max = array[index]
			maxI = index
		}
	}

	return maxI, max

}

func main() {

}
