package app

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helperSymmetric(root.Left, root.Right)
}

func helperSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && helperSymmetric(left.Right, right.Left) && helperSymmetric(left.Left, right.Right)
}
