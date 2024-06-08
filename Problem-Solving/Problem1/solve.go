package problem

import "fmt"

func FindDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	var count int
	for left < right {
		fmt.Println(left)
		fmt.Println(right)

		mid = left + (right-left)/2
		count = 0

		fmt.Println(mid)

		for num := range nums {
			if num <= mid {
				count = count + 1
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
