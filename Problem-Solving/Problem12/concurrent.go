package main

import (
	"fmt"
	"sort"
)

func topKFrequentConcurrent(nums []int, k int) []int {
	res := make([]int, k)
	freqMapChan := make(chan map[int]int)
	freqArrayChan := make(chan []Frequency)
	sortedChan := make(chan bool)

	go createFrequencyMap(nums, freqMapChan)
	go createFrequencyArray(freqArrayChan, freqMapChan)
	go createSortedArray(freqArrayChan, sortedChan)
	fmt.Println("waiting")
	<-sortedChan
	freqArray := <-freqArrayChan
	fmt.Println("creating results")
	for index := 0; index < k; index++ {
		res[index] = freqArray[index].key
	}
	fmt.Println("created results")
	close(freqMapChan)
	close(sortedChan)
	close(freqArrayChan)
	return res
}

func createFrequencyMap(nums []int, freqMapChan chan map[int]int) {
	fmt.Println("started creating map")
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	fmt.Println("done creating map")
	freqMapChan <- freqMap

}

func createFrequencyArray(freqArrayChan chan []Frequency, freqMapChan chan map[int]int) {
	freqMap := <-freqMapChan
	fmt.Println("Started creating array")
	var freqArray []Frequency
	for key, val := range freqMap {
		freqArray = append(freqArray, Frequency{key: key, val: val})
	}
	freqArrayChan <- freqArray
	fmt.Println("done creating array")
}

func createSortedArray(freqArrayChan chan []Frequency, sortedChan chan bool) {
	freqArray := <-freqArrayChan
	fmt.Println("started sorting")
	sort.Slice(freqArray, func(i, j int) bool {
		return freqArray[i].val > freqArray[j].val
	})
	fmt.Println("Finished sorting")
	sortedChan <- true
	freqArrayChan <- freqArray
}
