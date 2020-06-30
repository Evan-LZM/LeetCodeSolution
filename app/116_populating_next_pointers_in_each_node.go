package app

/*connect :Populating Next Right Pointers in Each Node
 You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.
*/

//solution.1.DFS
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

//2.BFS (using queue)
func connectBFS(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		m := len(queue)
		for i := 0; i < m; {
			h := queue[0]
			queue = queue[1:]
			i++
			if i == m {
				h.Next = nil
			} else {
				h.Next = queue[0]
			}

			if h.Left != nil {
				queue = append(queue, h.Left)
			}
			if h.Right != nil {
				queue = append(queue, h.Right)
			}
		}

	}
	return root
}
