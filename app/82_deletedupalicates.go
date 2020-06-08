package app

func DeleteSortedDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := skipDuplicates(head)
	if next == head.Next {
		head.Next = DeleteSortedDuplicates(head)
		return head
	}
	return DeleteSortedDuplicates(next)
}

func skipDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	if cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
		cur = cur.Next
	}
	return cur.Next
}
