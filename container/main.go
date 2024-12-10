package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MinHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	nums := []int{1, 4, 7, 6, 2, 3, 5}
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	fmt.Println(h)
	fmt.Println(heap.Pop(h))
	fmt.Println(h)
}
