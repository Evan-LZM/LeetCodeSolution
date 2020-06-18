package app

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		level := len(res)
		res = append(res, make([]int, len(queue)))
		newQueue := []*TreeNode{}
		for i, node := range queue {
			if level%2 == 0 {
				res[level][i] = node.Val
			} else {
				res[level][len(queue)-1-i] = node.Val
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return res
}
