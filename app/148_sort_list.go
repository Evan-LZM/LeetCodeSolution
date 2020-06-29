package app

/*SortList 题解：实现一个链表的归并排序即可。主要分为三部分：
快速排序
1.找到中点并返回的函数findMiddle;

2.归并函数merge;

3.排序函数sortList。
*/
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//now slow is the middle node for this list
	nextHalfHead := slow.Next
	slow.Next = nil
	//head---slow,nextHalfhead--end
	l1 := SortList(head)
	l2 := SortList(nextHalfHead)
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
