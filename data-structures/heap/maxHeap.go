package main

// MaxHeap represents a max heap.
type MaxHeap struct {
	data []int
}

// Insert adds an element to the heap and maintains the heap property.
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.upHeap(len(h.data) - 1)
}

// ExtractMax removes and returns the largest element (root) from the heap.
func (h *MaxHeap) ExtractMax() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	max := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.downHeap(0)
	return max
}

// upHeap maintains the heap property by bubbling up the element.
func (h *MaxHeap) upHeap(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] >= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// downHeap maintains the heap property by bubbling down the element.
func (h *MaxHeap) downHeap(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}
