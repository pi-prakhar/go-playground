package main

/**
PROBLEM:
leetcode -> partition -> medium

TAGS:
arrays, strings, recursion, backtracking

**/
func partition(s string) [][]string {
	var list [][]string
	solve(&s, &list, []string{}, 0)
	return list

}

func solve(s *string, list *[][]string, temp []string, start int) {
	if start == len(*s) {
		copyTemp := make([]string, len(temp))
		copy(copyTemp, temp)
		(*list) = append((*list), copyTemp)
	} else {
		for i := start; i < len(*s); i++ {
			if isPalindrome(s, start, i) {
				temp = append(temp, (*s)[start:i+1])
				solve(s, list, temp, i+1)
				temp = temp[:len(temp)-1]
			}
		}
	}
}

func isPalindrome(s *string, start int, end int) bool {
	for start <= end {
		if (*s)[start] != (*s)[end] {
			return false
		}
		start++
		end--
	}
	return true
}
