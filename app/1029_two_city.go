package app

import "sort"

func TwoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	result := 0
	for index, cost := range costs {
		if index < len(costs)/2 {
			result += cost[0]
		} else {
			result += cost[1]
		}
	}
	return result
}

func minCity(a, b int) int {
	if a > b {
		return b
	}
	return a
}
