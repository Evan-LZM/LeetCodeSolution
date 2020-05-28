package app

import "sort"

//PermuteUnique Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//1.recursion
func PermuteUnique2(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	temp := []int{}
	helpPermute2(nums, used, temp, &result)
	return result
}

func helpPermute2(nums []int, used []bool, temp []int, result *[][]int) {
	if len(nums) == len(temp) {
		r := append([]int{}, temp...)
		*result = append(*result, r)
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && !used[i-1] && nums[i] == nums[i-1]) {
			continue
		}
		used[i] = true
		temp = append(temp, nums[i])
		helpPermute2(nums, used, temp, result)
		used[i] = false
		temp = temp[:len(temp)-1]
	}
}
