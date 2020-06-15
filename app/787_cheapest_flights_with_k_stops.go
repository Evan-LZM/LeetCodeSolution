package app

import "math"

type node struct {
	key   int
	price int
}

//FindCheapestPrice by k stop
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][]node{}
	for _, flight := range flights {
		u, v, p := flight[0], flight[1], flight[2]
		graph[u] = append(graph[u], node{key: v, price: p})
	}
	queue := []node{{src, 0}}
	res := math.MaxInt32
	for len(queue) > 0 { 
		for size := len(queue); size > 0; size-- {
			u, c := queue[0].key, queue[0].price
			if u == dst {
				res = min(res, c)
			}
			for _, neigh := range graph[u] {
				v, p := neigh.key, neigh.price
				if c+p > res {
					continue
				}
				queue = append(queue, node{v, c + p})
			}
			queue = queue[1:]
		}
		if k > 0 {
			break
		}
		k--
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
