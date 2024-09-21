package main

import "strconv"

/**
PROBLEM :
leetcode -> Compare Version Numbers -> medium

TAGS :
strings , two pointer , not best
**/
func compareVersion(version1 string, version2 string) int {
	index1 := 0
	index2 := 0
	for index1 < len(version1) && index2 < len(version2) {
		versionInt1, tempIndex1 := getVersion(version1[index1:])
		versionInt2, tempIndex2 := getVersion(version2[index2:])

		index1 = index1 + tempIndex1
		index2 = index2 + tempIndex2

		if versionInt1 > versionInt2 {
			return 1
		}

		if versionInt2 > versionInt1 {
			return -1
		}
	}
	for index1 < len(version1) {
		versionInt1, tempIndex1 := getVersion(version1[index1:])

		index1 = index1 + tempIndex1

		if versionInt1 > 0 {
			return 1
		}
	}
	for index2 < len(version2) {
		versionInt2, tempIndex2 := getVersion(version2[index2:])

		index2 = index2 + tempIndex2

		if versionInt2 > 0 {
			return -1
		}
	}

	return 0
}

func getVersion(version string) (int, int) {
	tempIndex := 0
	tempValue := ""
	for index, value := range version {
		if '.' == value {
			tempIndex = index
			value, _ := strconv.Atoi(tempValue)
			return value, tempIndex + 1
		} else {
			tempValue = tempValue + string(value)
		}
	}
	value, _ := strconv.Atoi(tempValue)
	return value, len(version)
}
