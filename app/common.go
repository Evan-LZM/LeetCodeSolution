package app

type node struct {
	key   int
	price int
}

type KeyCoin struct {
	Amount    int
	CoinVlaue int
}


type city struct {
	visited bool
	name    string
}

type LRUCache struct {
	Head *LRUNode
	Tail *LRUNode
	HT   map[int]*LRUNode
	Cap  int
}

type LRUNode struct {
	Key  int
	Val  int
	Prev *LRUNode
	Next *LRUNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
	Random *Node
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	element []int
}