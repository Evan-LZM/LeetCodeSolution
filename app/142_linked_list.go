package app

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	entry := head
	for entry != slow {
		slow = slow.Next
		entry = entry.Next
	}
	return entry
}
