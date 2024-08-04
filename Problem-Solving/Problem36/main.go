package main

import "strconv"

func countAndSay(n int) string {
	if strconv.Itoa(n) == "1" {
		return "1"
	}
	return getRLEOfNumber(countAndSay(n - 1))
}

func getRLEOfNumber(s string) string {
	tempN := s[0]
	tempCount := 0
	retN := ""

	for i := 0; i < len(s); i++ {
		letter := s[i]
		if tempN == letter {
			tempCount++
		} else {
			retN = retN + strconv.Itoa(tempCount) + string(tempN)
			tempN = letter
			tempCount = 1
		}
	}

	retN = retN + strconv.Itoa(tempCount) + string(tempN)
	return retN
}
