package main

func containsDuplicate(nums []int) bool {
	map1 := make(map[int]int)
	for _, num := range nums {
		if _, ok := map1[num]; ok {
			return true
		}
		map1[num] = 1
	}
	return false
}
