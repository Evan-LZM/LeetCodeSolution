package app

/*LevelOrder solution:BFS (using queue)
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{} //traversal every root and add to this queue
	queue = append(queue, root)
	dummy := &TreeNode{}
	queue = append(queue, dummy)
	temp := []int{} //stor each layers number
	for len(queue) > 1 || len(temp) > 0 {
		h := queue[0]     //get the first node
		queue = queue[1:] //update queue
		if h == nil {
			continue
		}
		if h == dummy { //using dummy to distinguish diff layers
			l := append([]int{}, temp...)
			res = append(res, l)
			temp = []int{}
			queue = append(queue, dummy)
			continue
		}
		temp = append(temp, h.Val)
		if h.Left != nil {
			queue = append(queue, h.Left)
		}
		if h.Right != nil {
			queue = append(queue, h.Right)
		}
	}
	return res
}
