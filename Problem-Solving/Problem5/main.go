package main

import (
	"strings"
)

func isValidChar(r rune) bool {
	if (int(r) >= 97 && int(r) <= 122) || (int(r) >= 48 && int(r) <= 57) {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	len := len(s)
	indexEnd := len - 1
	indexStart := 0
	s = strings.ToLower(s)
	for indexStart <= indexEnd {
		if !isValidChar(rune(s[indexStart])) {
			indexStart++
			continue
		}

		if !isValidChar(rune(s[indexEnd])) {
			indexEnd--
			continue
		}

		if strings.Compare(string(s[indexStart]), string(s[indexEnd])) != 0 {
			return false
		}

		indexStart++
		indexEnd--

	}

	return true
}
