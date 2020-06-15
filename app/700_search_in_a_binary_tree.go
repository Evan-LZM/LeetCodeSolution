package app

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return rootmount_smbfs -d 777 -f 777 //learnazurefileshare27756:ZBGf4MvXr5wVfiTvCAGUedlEgrdQloacY7f4O4wtgWLUaM8GHLjBJ9zKGy25wAB82yePOO667MHYkvCHHP0Eag==@learnazurefileshare27756.file.core.windows.net learnazurefileshare27756
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
