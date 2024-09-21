package main

/**
PROBLEM:
leetcode -> Longest Palindromic Substring -> medium

TAGS:
strings, two pointer
**/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	length := len(s)
	for length > 0 {
		start := 0
		end := start + length
		for end <= len(s) {
			substring := s[start:end]
			if isPalindrome(substring) {
				return substring
			}
			start++
			end++
		}
		length--
	}
	return s
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
