package app

type BSTIterator struct { //stack
	stack []*TreeNode
}

func (this *BSTIterator) Next() int {
	cur := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	ret := cur.Val
	cur = cur.Right
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
	return ret
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
