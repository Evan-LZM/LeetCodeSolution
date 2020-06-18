package app

func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	index := searchInorder(postorder[len(postorder)-1], inorder)
	node := &TreeNode{
		Val:   inorder[index],
		Left:  buildTreeIP(inorder[:index], postorder[:index]),
		Right: buildTreeIP(inorder[index+1:], postorder[index:len(inorder)-1]),
	}
	return node
}

func searchInorder(val int, array []int) int {
	for i := 0; i < len(array); i++ {
		if val == array[i] {
			return i
		}
	}
	return 0
}
