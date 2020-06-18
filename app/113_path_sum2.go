package app

func solvePathSum2(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	helperPathSum2(root, sum, []int{}, &res)
	return res
}

func helperPathSum2(root *TreeNode, sum int, temp []int, res *[][]int) {
	//leaves
	if root.Left == nil && root.Right == nil {
		if root.Val == sum { //add temp to res
			cp := append([]int{}, temp...)
			cp = append(cp, root.Val)
			*res = append(*res, cp)
		}
		return
	}
	if root.Left != nil {
		temp = append(temp, root.Val)
		helperPathSum2(root.Left, sum-root.Val, temp, res)
		temp = temp[:len(temp)-1]
	}
	if root.Right != nil {
		temp = append(temp, root.Val)
		helperPathSum2(root.Right, sum-root.Val, temp, res)
		temp = temp[:len(temp)-1]
	}
}
