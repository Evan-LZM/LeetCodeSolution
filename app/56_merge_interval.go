package app

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	//first sort intervals Matrix array
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//res is the merged array
	res := [][]int{intervals[0]}
	lastValue := 0 //store the row index about intervals
	for i := 1; i < len(intervals); i++ {
		if res[lastValue][1] >= intervals[i][0] {
			res[lastValue][1] = max(res[lastValue][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			lastValue++
		}
	}
	return res
}

func maxMerge(x, y int) int {
	if x > y {
		return x
	}
	return y
}
