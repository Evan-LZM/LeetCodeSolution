package app

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	helperSum(root, &sum, 0)
	return sum
}

func helperSum(root *TreeNode, sum *int, temp int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*sum += 10*temp + root.Val
		return
	}
	if root.Left != nil {
		helperSum(root.Left, sum, 10*temp+root.Val)
	}

	if root.Right != nil {
		helperSum(root.Right, sum, 10*temp+root.Val)
	}
}
