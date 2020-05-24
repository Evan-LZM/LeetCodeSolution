package app

//MergeKLists a
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists == nil {
		return nil
	}
	return Khelper(lists, 0, len(lists)-1)
}

//Khelper a
func Khelper(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := (start + end) / 2
	a := Khelper(lists, start, mid)
	b := Khelper(lists, mid+1, end)
	return MergeTwoLists(a, b)
}
