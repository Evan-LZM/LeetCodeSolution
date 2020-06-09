package app

func ReverseBetween(head *ListNode, m, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	prev := dummy
	mprev := dummy
	mcur := cur
	ncur := cur
	nlast := cur
	for i := 1; i <= n; i++ {
		next := cur.Next
		if i == m-1 {
			mprev = cur
		}
		if i == m {
			mcur = cur
		}
		if i == n {
			ncur = cur
			nlast = cur.Next
		}
		if i >= m+1 && i <= n {
			cur.Next = prev
		}
		prev = cur
		cur = next
	}
	mprev.Next = ncur
	mcur.Next = nlast
	return dummy.Next
}
