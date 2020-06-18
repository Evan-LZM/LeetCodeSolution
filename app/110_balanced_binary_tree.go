package app

func isBalanced(root *TreeNode) bool {
	if getDepth(root) < 0 {
		return false
	}
	return true
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	if left == -1 {
		return -1
	}
	right := getDepth(root.Right)
	if right == -1 {
		return -1
	}
	diff := left - right
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return -1
	}
	return max(left, right) + 1
}
