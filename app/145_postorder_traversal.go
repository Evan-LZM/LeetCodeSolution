package app

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	helperPostOrder(root, &res)
	return res
}

func helperPostOrder(root *TreeNode, sum *[]int) {
	if root != nil {
		helperPostOrder(root.Left, sum)
		helperPostOrder(root.Right, sum)
		*sum = append(*sum, root.Val)
		return
	}
}
