package main

import (
	"testing"
)

func isEqual(got *ListNode, want *ListNode) bool {
	if got == nil && want == nil {
		return true
	}
	if got.Val != want.Val {
		return false
	}
	return true && isEqual(got.Next, want.Next)
}

func Test_MergeTwoLists(t *testing.T) {
	tests := []struct {
		testName string
		list1    *ListNode
		list2    *ListNode
		want     *ListNode
	}{
		{
			"Multiple Nodes",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			"One Node",
			nil,
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			"No Node",
			nil,
			nil,
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := mergeTwoLists(test.list1, test.list2)
			if !isEqual(got, test.want) {
				t.Errorf("Failed %s: expected %+v, got %+v", test.testName, test.want, got)
			}
		})
	}
}
