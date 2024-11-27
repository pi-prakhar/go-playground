package main

import "fmt"

func main() {
	var data []int

	minHeap := MinHeap{
		data: data,
	}

	minHeap.Insert(7)
	minHeap.Insert(5)
	minHeap.Insert(3)
	minHeap.Insert(2)
	minHeap.Insert(4)
	minHeap.Insert(1)

	fmt.Println(minHeap.data)
	fmt.Printf("Min Value : %d \n", minHeap.ExtractMin())
	fmt.Println(minHeap.data)
}
