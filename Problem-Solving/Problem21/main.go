package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func copyLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val}
	currentOriginal := head.Next
	currentNew := newHead

	for currentOriginal != nil {
		newNode := &ListNode{Val: currentOriginal.Val}
		currentNew.Next = newNode
		currentNew = newNode
		currentOriginal = currentOriginal.Next
	}

	return newHead
}
func getReverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	temp := getReverseList(head.Next)

	temp2 := head.Next
	head.Next = head.Next.Next
	temp2.Next = head

	return temp

}
func reorderList(head *ListNode) {
	list := copyLinkedList(head)
	reverseList := getReverseList(head)

	temp := &ListNode{
		Val: list.Val, Next: nil,
	}
	tempList := temp
	list = list.Next

	alternate := true

	for list != reverseList {
		if alternate {
			temp = &ListNode{
				Val: list.Val, Next: nil,
			}
			tempList.Next = temp
			reverseList = reverseList.Next
			alternate = false
		} else {
			temp = &ListNode{
				Val: list.Val, Next: nil,
			}
			tempList.Next = temp
			list = list.Next
			alternate = true
		}
	}

	head = tempList
}
