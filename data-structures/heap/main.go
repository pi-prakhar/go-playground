package main

import "fmt"

func main() {
	var dataMin []int
	var dataMax []int

	minHeap := MinHeap{
		data: dataMin,
	}

	maxHeap := MaxHeap{
		data: dataMax,
	}

	minHeap.Insert(7)
	minHeap.Insert(5)
	minHeap.Insert(3)
	minHeap.Insert(2)
	minHeap.Insert(4)
	minHeap.Insert(1)

	fmt.Println("Min Heap before : ", minHeap.data)
	fmt.Printf("Min Value : %d \n", minHeap.ExtractMin())
	fmt.Println("Min Heap after : ", minHeap.data)

	maxHeap.Insert(7)
	maxHeap.Insert(5)
	maxHeap.Insert(3)
	maxHeap.Insert(2)
	maxHeap.Insert(4)
	maxHeap.Insert(1)

	fmt.Println("Max Heap before : ", maxHeap.data)
	fmt.Printf("Max Value : %d \n", maxHeap.ExtractMax())
	fmt.Println("Max Heap after : ", maxHeap.data)
}
