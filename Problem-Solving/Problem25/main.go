package main

/**
PROBLEM:
leetcode -> permute -> medium

Tags:
recursion, arrays, backtracking
**/

func permute(nums []int) [][]int {
	var list [][]int
	// sort.Slice(nums, func(i int, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	set := make(map[int]struct{})
	solve(&list, &nums, []int{}, &set)
	return list

}

func solve(list *[][]int, nums *[]int, templist []int, set *map[int]struct{}) {
	if len(templist) == len(*nums) {
		copyOfTemp := make([]int, len(templist))
		copy(copyOfTemp, templist)
		(*list) = append((*list), copyOfTemp)
	} else {
		for i := 0; i < len((*nums)); i++ {
			if _, ok := (*set)[(*nums)[i]]; ok {
				continue
			} else {
				templist = append(templist, (*nums)[i])
				(*set)[(*nums)[i]] = struct{}{}
				solve(list, nums, templist, set)
				delete((*set), templist[len(templist)-1])
				templist = templist[:len(templist)-1]
			}

		}

	}

}
