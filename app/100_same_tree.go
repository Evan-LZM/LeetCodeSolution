package app

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}
	return q.Val == p.Val && IsSameTree(q.Left, p.Left) && IsSameTree(q.Right, p.Right)
}
