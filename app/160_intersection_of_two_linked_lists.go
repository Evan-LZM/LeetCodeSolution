package app

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cura, curb := headA, headB
	for cura != nil && curb != nil {
		cura = cura.Next
		curb = curb.Next
	}
	for cura != nil {
		cura = cura.Next
		headA = headA.Next
	}
	for curb != nil {
		curb = curb.Next
		headB = headB.Next
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
