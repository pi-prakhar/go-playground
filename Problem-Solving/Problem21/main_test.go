package main

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice of integers
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice of integers
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Test function for reorderList using the testing library
func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4}, expected: []int{1, 4, 2, 3}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 5, 2, 4, 3}},
		{input: []int{1}, expected: []int{1}},
	}

	for _, test := range tests {
		head := createLinkedList(test.input)
		reorderList(head)
		result := linkedListToSlice(head)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
