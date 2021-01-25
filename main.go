package main

import (
	"fmt"
)

type Node struct {
	val      int
	children []Node
}

type Data struct {
	list []Node
}

func main() {
    if 4<4{
        fmt.Println("Yes")
    }else{
        fmt.Println("No")
    }
	a := Node{val: 1}
	b := Node{val: 2}
	b.children = append(b.children, a)
	c := Node{val: 3}
	c.children = append(c.children, b)
	e := Node{val: 4}
	f := Node{val: 4}
	g := Node{val: 4}
	h := Node{val: 4}
	f.children = append(f.children, e)
	g.children = append(g.children, f)
	h.children = append(h.children, g)
	d := Node{val: 4}
	d.children = append(d.children, h)
	d.children = append(d.children, c)
	result := Data{}
	result.list = append(result.list, d)
	fmt.Println(len(result.list))
	fmt.Println(maxDepth(result.list))
}

func maxDepth(list []Node) int {
	if len(list) == 0 {
		return 0
	}
	max := 0
	for _, v := range list {
		dep := maxDepth(v.children)
		if max < dep {
			max = dep
		}
	}
	return max + 1
}
