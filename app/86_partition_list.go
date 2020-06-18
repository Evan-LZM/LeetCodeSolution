package app

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	//143252
	//seprate into two list
	//1:122
	//2:435
	//integrate:122435
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	//find first node which larger than x
	cur1 := dummy1
	cur2 := dummy2
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur1.Next
		} else {
			cur2.Next = cur
			cur2 = cur2.Next
		}
	}
	cur1.Next = dummy2.Next
	cur2.Next = nil
	return dummy1.Next
}
