package main

/**
PROBLEM:
Leetcode -> Longest Substring Without Repeating Characters -> medium

TAGS:
hashtable, strings, two pointer, sliding window

**/

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := make(map[rune]int)
	left := 0
	for right, character := range s {
		if index, ok := m[character]; ok && index >= left {
			left = index + 1
		} else if maxLen < (right - left + 1) {
			maxLen = right - left + 1
		}
		m[character] = right
	}
	return maxLen
}

// func lengthOfLongestSubstring(s string) int {
// 	maxLen := 0
// 	currLen := 0
// 	substring := ""
// 	for _, character := range s {
// 		if strings.ContainsRune(substring, character) {
// 			index := strings.Index(substring, string(character))
// 			substring = substring[index+1:]
// 		}
// 		substring = substring + string(character)
// 		currLen = len(substring)
// 		if currLen > maxLen {
// 			maxLen = currLen
// 		}
// 	}
// 	return maxLen
// }
