package app

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	subset := make([]int, 0)
	sort.Ints(nums)
	helperSubsetsWithDup(nums, &result, subset, 0)
	return result
}

func helperSubsetsWithDup(nums []int, data *[][]int, subset []int, index int) {
	temp := make([]int, len(subset))
	copy(temp, subset)
	*data = append(*data, temp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		helperSubsetsWithDup(nums, data, subset, i+1)
		subset = subset[:len(subset)-1]
	}
}
