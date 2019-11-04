package app

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	result := 0
	sort.Ints(nums)
	diff := math.MaxInt32
	for i := 0; i <= len(nums)-3; i++ {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if getAbsolute(sum-target) < diff {
				result = sum
				diff = getAbsolute(sum - target)
			}
			if sum == target {
				return sum
			} else if sum-target > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
func getAbsolute(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
