package app


func Constructor(capacity int) LRUCache {
	return LRUCache{HT: make(map[int]*LRUNode, capacity), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.HT[key]
	if ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.HT[key]
	if ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	}
	node = &LRUNode{Key: key, Val: value}
	this.HT[key] = node
	this.Add(node)
	if len(this.HT) > this.Cap {
		delete(this.HT, this.Tail.Key)
		this.Remove(this.Tail)
	}
}

//插入头
func (this *LRUCache) Add(node *LRUNode) {
	node.Prev = nil
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) Remove(node *LRUNode) {
	if node != this.Head {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}
	if node != this.Tail {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}
}
