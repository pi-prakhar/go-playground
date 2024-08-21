package main

/**
PROBLEM:
leetcode -> Find the Index of the First Occurrence in a String -> easy

TAGS:
string
**/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			if checkMatch(haystack, needle, i) {
				return i
			}
		}
	}

	return -1
}

func checkMatch(haystack, needle string, index int) bool {
	for i := 1; i < len(needle); i++ {
		if haystack[index+i] != needle[i] {
			return false
		}
	}
	return true
}
