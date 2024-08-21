package main

func strStr(haystack string, needle string) int {
	for index, _ := range haystack {
		if haystack[index] == needle[0] && check(haystack, needle, index) {
			return index
		}
	}

	return -1
}

func check(haystack string, needle string, index int) bool {
	for i, _ := range needle {
		if index >= len(haystack) || needle[i] != haystack[index] {
			return false
		}
		index++
	}
	return true
}
