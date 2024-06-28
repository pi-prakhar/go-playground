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

func Test_ReorderList(t *testing.T) {
	tests := []struct {
		testName string
		list     *ListNode
		want     *ListNode
	}{
		{
			"Multiple Nodes",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			&ListNode{1, &ListNode{5, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}},
		},
		{
			"Two Nodes",
			&ListNode{1, &ListNode{2, nil}},
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			"One Node",
			&ListNode{1, nil},
			&ListNode{1, nil},
		},
		{
			"No Node",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			&ListNode{1, &ListNode{4, &ListNode{2, &ListNode{3, nil}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := reorderList(test.list)
			if !isEqual(got, test.want) {
				t.Errorf("Failed %s: expected %+v, got %+v", test.testName, test.want, got)
			}
		})
	}
}
