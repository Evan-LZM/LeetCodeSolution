package app

//data:2,3,6,7 target 7
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) == 0 {
		return result
	}
	temp := []int{}
	helpCombinationSum(candidates, temp, target, &result, 0)
	return result
}

func helpCombinationSum(candidates []int, temp []int, target int, result *[][]int, p int) {
	if p == len(candidates) || target < 0 {
		return
	}
	if target == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
	}
	for i := p; i < len(candidates); i++ {
		t := append(temp, candidates[i])
		helpCombinationSum(candidates, t, target-candidates[i], result, i)
	}
}
