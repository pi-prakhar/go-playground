package main

//Two Sum
func twoSum(nums []int, target int) []int {
	remainderMap := make(map[int]int)
	result := make([]int, 2)
	for index, num := range nums {
		remainder := target - num
		if val, ok := remainderMap[remainder]; ok {
			result[0] = val
			result[1] = index
		} else {
			remainderMap[num] = index
		}
	}
	return result
}
