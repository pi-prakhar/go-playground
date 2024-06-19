package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longestStreak := 0

	for num := range numSet {
		// Only start counting when `num` is the start of a sequence
		if _, found := numSet[num-1]; !found {
			currentNum := num
			currentStreak := 1

			for {
				if _, found := numSet[currentNum+1]; found {
					currentNum += 1
					currentStreak += 1
				} else {
					break
				}
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Longest consecutive sequence length:", longestConsecutive(nums)) // Output: 4 (sequence: 1, 2, 3, 4)
}
