package main

/**
PROBLEM:
leetcode -> Reorder List -> medium

TAGS:
linkedlist ->linked list-> double pointer
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Rearrange reorders the linked list as specified
func reorderList(head *ListNode) {
	rearrange(head, head)
}

// // Rearrange reorders the linked list as specified
// func rearrange(head **ListNode, last *ListNode) {
// 	if last == nil {
// 		return
// 	}

// 	// Recursive call
// 	rearrange(head, last.Next)

// 	// (*head).Next will be set to nil after rearrangement
// 	if (*head).Next == nil {
// 		return
// 	}

// 	// Rearrange the list until both head and last meet or next to each other
// 	if *head != last && (*head).Next != last {
// 		tmp := (*head).Next
// 		(*head).Next = last
// 		last.Next = tmp
// 		*head = tmp
// 	} else {
// 		if *head != last {
// 			*head = (*head).Next
// 		}
// 		(*head).Next = nil
// 	}
// }

func rearrange(head *ListNode, last *ListNode) {
	if last == nil {
		return
	}

	rearrange(head, last.Next) // This modifies a copy of head, not the original head pointer

	if head.Next == nil {
		return
	}

	if head != last && head.Next != last {
		tmp := head.Next
		head.Next = last
		last.Next = tmp
		head = tmp // This modifies a copy of head, not the original head pointer
	} else {
		if head != last {
			head = head.Next // This modifies a copy of head, not the original head pointer
		}
		head.Next = nil // This modifies a copy of head, not the original head pointer
	}
}
