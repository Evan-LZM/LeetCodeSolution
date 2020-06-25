package app

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	helperPreOrder(root, &result)
	return result
}

func helperPreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	helperPreOrder(root.Left, res)
	helperPreOrder(root.Right, res)
}
