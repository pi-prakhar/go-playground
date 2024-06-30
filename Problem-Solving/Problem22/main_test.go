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

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		testName string
		head     *ListNode
		n        int
		want     *ListNode
	}{
		{
			"Multiple Nodes",
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			"Two Nodes",
			&ListNode{1, &ListNode{2, nil}},
			1,
			&ListNode{1, nil},
		},
		{
			"One Node",
			&ListNode{1, nil},
			1,
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := removeNthFromEnd(test.head, test.n)
			if !isEqual(got, test.want) {
				t.Errorf("Failed %s: expected %+v, got %+v", test.testName, test.want, got)
			}
		})
	}
}
