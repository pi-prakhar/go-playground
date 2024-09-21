package main

import (
	"testing"
)

func addCycle(list *ListNode, start int, end int) {
	index := 0
	temp := list
	var startNode *ListNode
	for list != nil {
		if index == start {
			startNode = list
		}
		if index == end {
			list.Next = startNode
			break
		}
		list = list.Next
		index++
	}
	list = temp
}

func Test_MergeTwoLists(t *testing.T) {
	tests := []struct {
		testName string
		list     *ListNode
		want     bool
	}{
		{
			"Multiple Nodes has cycle",
			&ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}},
			true,
		},
		{
			"Multiple Nodes No cycle",
			&ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}},
			false,
		},
		{
			"One Node",
			&ListNode{1, nil},
			false,
		},
		{
			"No Node",
			nil,
			false,
		},
	}
	//add cycle in list1
	addCycle(tests[0].list, 1, 4)

	//testing now
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := hasCycle(test.list)
			if got != test.want {
				t.Errorf("Failed %s: expected %t, got %t", test.testName, test.want, got)
			}
		})
	}
}
