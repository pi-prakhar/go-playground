package main

import "sort"

/**
Problem
leetcode -> Group Anagrams -> medium

Tags
hashtable, strings, sorting, arrays
**/

func getSortedString(s *string) string {
	runes := []rune(*s)
	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert the sorted slice of runes back to a string
	return string(runes)

}
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	//create map to store values
	store := make(map[string][]string)

	//iterate over strs
	for _, str := range strs {
		//get sorted str
		sortedStr := getSortedString(&str)

		//check if it is in the store
		if group, ok := store[sortedStr]; ok {
			//modify the group
			group = append(group, str)
			store[sortedStr] = group

		} else {
			//create a new group
			store[sortedStr] = []string{str}
		}
	}

	for _, val := range store {
		result = append(result, val)
	}

	return result
}

// func getSortedString(s *string) string {
// 	runes := []rune(*s)
// 	// Sort the slice of runes
// 	sort.Slice(runes, func(i, j int) bool {
// 		return runes[i] < runes[j]
// 	})

// 	// Convert the sorted slice of runes back to a string
// 	return string(runes)

// }
// func groupAnagrams(strs []string) [][]string {
// 	var result [][]string
// 	//create map to store values
// 	store := make(map[int]map[string][]string)

// 	//iterate over strs
// 	for _, str := range strs {
// 		//get sorted str
// 		sortedStr := getSortedString(&str)

// 		//check if it is in the store
// 		if groupCount, ok := store[len(sortedStr)]; ok {
// 			//modify the group
// 			if group, ok := groupCount[sortedStr]; ok {
// 				group = append(group, str)
// 				groupCount[sortedStr] = group
// 			} else {
// 				groupCount[sortedStr] = []string{str}
// 			}

// 		} else {
// 			//create a new group
// 			newGroup := make(map[string][]string)
// 			newGroup[sortedStr] = []string{str}
// 			store[len(sortedStr)] = newGroup
// 		}
// 	}

// 	for _, groups := range store {
// 		for _, val := range groups {
// 			result = append(result, val)
// 		}
// 	}

// 	return result
// }
