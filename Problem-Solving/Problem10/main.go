package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func checkCycle(head1 *ListNode, head2 *ListNode) bool {
	if head1 == nil || head2 == nil || head2.Next == nil || head2.Next.Next == nil {
		return false
	}

	if head1 == head2 {
		return true
	}

	return checkCycle(head1.Next, head2.Next.Next)
}
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	return checkCycle(head, head.Next)
}
