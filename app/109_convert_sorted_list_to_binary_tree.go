package app

func sortedListToBST(head *ListNode) *TreeNode {
	return helperBST(head, nil)
}

func helperBST(head *ListNode, end *ListNode) *TreeNode {
	if head == end {
		return nil
	}
	if head.Next == end {
		return &TreeNode{Val: head.Val}
	}
	mid := findMid(head, end)
	return &TreeNode{
		Val:   mid.Val,
		Left:  helperBST(head, mid),
		Right: helperBST(mid.Next, end),
	}
}

func findMid(start *ListNode, end *ListNode) *ListNode {
	fast := start //double speed than slow==*2
	slow := start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
