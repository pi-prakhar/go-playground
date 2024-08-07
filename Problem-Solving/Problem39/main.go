package main

func nextPermutation(nums []int) {
	length := len(nums)
	startBehind := -1
	for last := length - 2; last >= 0; last-- {
		if nums[last] < nums[last+1] {
			start := last + 1
			target := nums[last]
			end := larger(&nums, &target, &start)
			swap(&nums, &last, &end)
			reverseNums(&nums, &last, &length)
			return
		}
	}

	reverseNums(&nums, &startBehind, &length)
	return
}

func reverseNums(nums *[]int, start *int, end *int) {
	*start++
	*end--
	for *start <= *end {
		swap(nums, start, end)
		*start++
		*end--
	}
}

func swap(nums *[]int, start *int, end *int) {
	temp := (*nums)[*start]
	(*nums)[*start] = (*nums)[*end]
	(*nums)[*end] = temp
}

func larger(nums *[]int, target *int, start *int) int {
	minV := (*nums)[*start] - (*target)
	minI := *start
	for end := *start; end < len(*nums); end++ {
		temp := (*nums)[end] - (*target)
		if temp > 0 && temp <= minV {
			minV = temp
			minI = end
		}
	}
	return minI
}
