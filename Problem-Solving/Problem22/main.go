package main

/**
PROBLEM:
leetcode -> Remove Nth Node From End of List -> medium

TAGS:
linkedlist ->linked list-> double pointer
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head **ListNode, count *int) *ListNode {
	if (*head).Next == nil {
		(*count)--
		if (*count) == 0 {
			return nil
		}
		return (*head)
	}

	(*head).Next = solve(&(*head).Next, count)
	(*count)--
	if *count == 0 {
		return (*head).Next
	}
	return *head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return solve(&head, &n)
}
