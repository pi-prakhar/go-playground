package main

// MinHeap represents a min heap.
type MinHeap struct {
	data []int
}

// Insert adds an element to the heap and maintains the heap property.
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.upHeap(len(h.data) - 1)
}

// ExtractMin removes and returns the smallest element (root) from the heap.
func (h *MinHeap) ExtractMin() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.downHeap(0)
	return min
}

// upHeap maintains the heap property by bubbling up the element.
func (h *MinHeap) upHeap(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] <= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// downHeap maintains the heap property by bubbling down the element.
func (h *MinHeap) downHeap(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}
