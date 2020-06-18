package app

//反方向插入可以换一个思路。即永远插入第一个位置。之后将第一个位置转移到第二个位置。第一个位置保持空
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		temp := append([][]int{}, []int{})
		result = append(temp, result...)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			result[0] = append(result[0], node.Val)
		}
		queue = queue[size:]
	}
	return result
}
