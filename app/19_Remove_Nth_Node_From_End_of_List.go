package app

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	for ; n > 0; n-- {
		head = head.Next
	}
	prev := dummy //pass the address
	for head != nil {
		head, prev = head.Next, prev.Next
	}
	prev.Next = prev.Next.Next
	for dummy != nil {
		if dummy != nil {
			dummy = dummy.Next
		}
	}
	return dummy.Next
}
