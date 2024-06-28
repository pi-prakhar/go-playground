package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func copyLinkedList(head *ListNode) (*ListNode, int) {
	if head == nil {
		return nil, -1
	}

	newHead := &ListNode{Val: head.Val}
	currentOriginal := head.Next
	currentNew := newHead
	count := 1

	for currentOriginal != nil {
		count++
		newNode := &ListNode{Val: currentOriginal.Val}
		currentNew.Next = newNode
		currentNew = newNode
		currentOriginal = currentOriginal.Next
	}

	return newHead, count
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
func reorderList(head *ListNode) *ListNode {
	list, count := copyLinkedList(head)
	reverseList := getReverseList(head)

	temp := &ListNode{
		Val: list.Val, Next: nil,
	}
	tempList := temp
	list = list.Next
	count--

	alternate := true
	head = tempList
	for count > 0 {
		if alternate {
			temp = &ListNode{
				Val: reverseList.Val, Next: nil,
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
		count--
		tempList = tempList.Next
	}
	tempList = nil
	return head
}
