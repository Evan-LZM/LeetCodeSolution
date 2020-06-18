package app

func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return helpPathSum(root, sum)
}

func helpPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	}
	//root!=nil:1.左右都为空，2.左为空。3.右为空.4.左右都不为空
	//左右都为空代表为叶子节点
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	//左为空。则不是叶子节点
	if root.Left == nil {
		return helpPathSum(root.Right, sum-root.Val)
	}
	//3.右为空.
	if root.Right == nil {
		return helpPathSum(root.Left, sum-root.Val)
	}
	//4.左右都不为空
	return helpPathSum(root.Right, sum-root.Val) || helpPathSum(root.Left, sum-root.Val)
}
