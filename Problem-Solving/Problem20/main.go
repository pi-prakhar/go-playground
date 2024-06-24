package main

func findMin(nums []int) int {
	l := len(nums)
	left := 0
	right := l - 1
	mid := (left + right) / 2
	gap := mid - left

	for !((nums[left] <= nums[mid]) && (nums[mid] <= nums[right])) {
		if (nums[mid] >= nums[left]) && (nums[left] >= nums[right]) {
			left = right
		} else {
			left = mid
		}
		right = (left + l - 1) % l
		mid = (left + gap) % l
	}
	return left
}

func search(nums []int, target int) int {
	tl := len(nums)
	left := findMin(nums)
	l := tl

	for l >= 0 {
		mid := (left + (l / 2)) % tl
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			l = l / 2
		} else {
			left = (mid + 1) % tl
			l = l / 2
		}

	}
	return -1
}
