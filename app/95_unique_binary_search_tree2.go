package app

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var nodes []int
	for i := 1; i < n+1; i++ {
		nodes = append(nodes, i)
	}
	return buildTree(nodes)
}

func buildTree(nodes []int) []*TreeNode {
	if len(nodes) == 0 {
		return []*TreeNode{nil}
	}
	var ret []*TreeNode
	for i := 0; i < len(nodes); i++ {
		lt := buildTree(nodes[:i])
		rt := buildTree(nodes[i+1:])
		for _, l := range lt {
			for _, r := range rt {
				ret = append(ret, &TreeNode{
					Val:   nodes[i],
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return ret
}
