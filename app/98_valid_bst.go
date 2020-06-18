package app

/*IsValidBST Validate Binary Search Tree
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/
func IsValidBST(root *TreeNode) bool {
	ret, _, _ := helperValidBST(root)
	return ret
}

func helperValidBST(cur *TreeNode) (bool, int, int) {
	if cur == nil {
		return true, 0, 0
	}
	//node is leaves
	if cur.Left == nil && cur.Right == nil {
		return true, cur.Val, cur.Val
	}
	leftValid, leftmin, leftmax := helperValidBST(cur.Left)
	if !leftValid {
		return false, 0, 0
	}
	rightValid, rightmin, rightmax := helperValidBST(cur.Right)
	if !rightValid {
		return false, 0, 0
	}
	//left is nil
	if cur.Left == nil && cur.Right != nil {
		if cur.Val < rightmin {
			return true, cur.Val, rightmax
		}
		return false, 0, 0
	}
	//right is nil
	if cur.Left != nil && cur.Right == nil {
		if cur.Val > leftmax {
			return true, leftmin, cur.Val
		}
		return false, 0, 0
	}
	//bot are not nil
	if leftmax < cur.Val && cur.Val < rightmin {
		return true, leftmin, rightmax
	}
	return false, 0, 0
}
