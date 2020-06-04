package app

func Subsets(nums []int) [][]int {
	return helperSubsets(nums, 0)
}

func helperSubsets(nums []int, p int) [][]int {
	var result [][]int
	result = append(result, []int{})
	if p == len(nums) {
		return result
	}
	for i := p; i < len(nums); i++ {
		temp := helperSubsets(nums, i+1)
		for j := range temp {
			t := append([]int{nums[i]}, temp[j]...)
			result = append(result, t)
		}
	}
	return result
}
