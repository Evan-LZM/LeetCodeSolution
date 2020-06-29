package app

/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

这道链表重排序问题可以拆分为以下三个小问题：

1. 使用快慢指针来找到链表的中点，并将链表从中点处断开，形成两个独立的链表。

2. 将第二个链翻转。

3. 将第二个链表的元素间隔地插入第一个链表中。
*/
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodeList []*ListNode
	for cur := head; cur != nil; cur = cur.Next {
		nodeList = append(nodeList, cur)
	}
	i := 0
	j := len(nodeList) - 1
	//seperate to two list
	nodeList[(len(nodeList)-1)/2].Next = nil
	for i < j {
		nodei := nodeList[i]
		nodej := nodeList[j]
		nodej.Next = nodei.Next
		nodei.Next = nodej
		i++
		j--
	}
}
