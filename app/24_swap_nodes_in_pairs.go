package app

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		temp := cur.Next.Next
		cur.Next.Next = cur
		cur.Next = temp
		prev = cur
		cur = temp
	}
	return dummy.Next
}
