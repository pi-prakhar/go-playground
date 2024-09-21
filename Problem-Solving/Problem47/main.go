package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsStr := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		numsStr[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(numsStr, func(i, j int) bool {
		return numsStr[i]+numsStr[j] > numsStr[j]+numsStr[i]
	})

	if numsStr[0] == "0" {
		return "0"
	}

	return strings.Join(numsStr, "")
}
