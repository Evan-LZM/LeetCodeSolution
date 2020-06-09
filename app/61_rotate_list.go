package app

/*RotateRight Given a linked list, rotate the list to the right by
k places, where k is non-negative.
Example 1:
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
*/
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	len := 0
	for cur != nil {
		cur = cur.Next
		len++
	}
	k = k % len
	if k == 0 {
		return head
	}
	cur = head
	prev := dummy
	temp := head
	last := dummy
	for i := 0; i < k; i++ {
		temp = temp.Next
		last = last.Next
	}
	for temp != nil {
		temp = temp.Next
		last = last.Next
		prev = prev.Next
		cur = cur.Next
	}
	last.Next = head
	prev.Next = nil
	return cur
}
