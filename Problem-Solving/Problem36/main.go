package main

import (
	"bytes"
	"strconv"
)

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

func countAndSayOptimized(n int) string {
	if n == 1 {
		return "1"
	}

	prev := []byte("1")
	for i := 2; i <= n; i++ {
		prev = getRLEOfNumberOptimized(prev)
	}

	return string(prev)
}

func getRLEOfNumberOptimized(s []byte) []byte {
	var result bytes.Buffer
	result.Grow(len(s) * 2) // Preallocate memory

	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(s[i-1])
			count = 1
		}
	}

	result.WriteString(strconv.Itoa(count))
	result.WriteByte(s[len(s)-1])

	return result.Bytes()
}
