package app

import (
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	temp := []int{}
	helpCombinationSum2(candidates, temp, target, &result, -1)
	return result
}

func helpCombinationSum2(candidates []int, temp []int, target int, result *[][]int, p int) {
	if target < 0 {
		return
	}
	if target == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}
	for i := p + 1; i < len(candidates); i++ {
		if i != p+1 && candidates[i] == candidates[i-1] {
			continue
		}
		t := append(temp, candidates[i])
		helpCombinationSum2(candidates, t, target-candidates[i], result, i)
	}
}
