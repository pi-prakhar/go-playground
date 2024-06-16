package main

import (
	"sort"
)

type Frequency struct {
	key int
	val int
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	res := make([]int, k)
	for _, num := range nums {
		if count, ok := freqMap[num]; ok {
			count++
			freqMap[num] = count
		} else {
			freqMap[num] = 1
		}
	}
	var freqArray []Frequency
	for key, val := range freqMap {
		freqArray = append(freqArray, Frequency{key: key, val: val})
	}

	sort.Slice(freqArray, func(i, j int) bool {
		return freqArray[i].val > freqArray[j].val
	})

	for index := 0; index < k; index++ {
		res[index] = freqArray[index].key
	}

	return res

}
