package main

/**
PROBLEM:
leetcode -> Longest Repeating Character Replacement  -> medium

Tags:
remember , algorithm, arrays, two pointer, hashtable
**/
func characterReplacement(s string, k int) int {
	store := make(map[byte]int)
	res := 0
	max := 0

	left := 0
	for right := 0; right < len(s); right++ {
		currCharRight := s[right]
		store[currCharRight]++

		if store[currCharRight] > max {
			max = store[currCharRight]
		}

		if right-left+1-max > k {
			currCharLeft := s[left]
			store[currCharLeft]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
