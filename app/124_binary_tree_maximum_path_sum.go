package app

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := root.Val
	helperMaxPathSum(root, &result)
	return result
}

func helperMaxPathSum(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := helperMaxPathSum(root.Left, result)
	right := helperMaxPathSum(root.Right, result)
	*result = max(*result, left+right+root.Val)
	lmax := left + root.Val
	rmax := right + root.Val
	return max(lmax, rmax)
}
