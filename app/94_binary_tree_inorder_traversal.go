package app

/*InorderTraversal
 * Definition for a binary tree node.
Inorder :中序遍历==》左根右
*/

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	InorderRecursion(root, &result)
	return result
}
func InorderRecursion(root *TreeNode, temp *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		InorderRecursion(root.Left, temp)
	}
	if root != nil {
		*temp = append(*temp, root.Val)
	}
	if root.Right != nil {
		InorderRecursion(root.Right, temp)
	}
}
